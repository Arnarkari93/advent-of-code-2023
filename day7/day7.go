package main

import (
	"advent_of_code_2023/utils"
	"fmt"
	"slices"
	"sort"
	"strings"
	"time"
)

func main() {
	filePath := utils.GetFilePathFromArgs()
	input := utils.ReadFileToArray(filePath, "\n")
	input = input[:len(input)-1]
	start := time.Now()
	Part1(input)
	fmt.Println("Part 1 took", time.Since(start))

	start = time.Now()
	Part2(input)
	fmt.Println("Part 2 took", time.Since(start))
}

func Part1(input []string) {
	hands := parseInput(input)

	var cardRanking = map[byte]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': 11,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
	}
	sort.Slice(hands, func(a, b int) bool {
		aComp := HandAndStrength{
			strength: GetStrengthOfHand(hands[a].hand),
			hand:     hands[a].hand,
		}
		bComp := HandAndStrength{
			strength: GetStrengthOfHand(hands[b].hand),
			hand:     hands[b].hand,
		}
		return handComp(aComp, bComp, cardRanking)
	})

	sum := sumOfBids(hands)

	// fmt.Println(hands)
	fmt.Println(sum)
}

func Part2(input []string) {
	hands := parseInput(input)

	var cardRanking = map[byte]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
		'J': 1,
	}

  cache := map[string]HandStrength{}
	sort.Slice(hands, func(a, b int) bool {
		aComp := HandAndStrength{
			strength: GetStrengthOfHandWithJokers(hands[a].hand, cardRanking, &cache),
			hand:     hands[a].hand,
		}
		bComp := HandAndStrength{
			strength: GetStrengthOfHandWithJokers(hands[b].hand, cardRanking, &cache),
			hand:     hands[b].hand,
		}
		return handComp(aComp, bComp, cardRanking)
	})

	sum := sumOfBids(hands)

	// fmt.Println(hands)
	fmt.Println(sum)
}
func ParseLine(line string) (string, int) {
	lineSplit := strings.Fields(line)
	hand, bid := lineSplit[0], utils.DangerouslyParseInt(lineSplit[1])
	return hand, bid
}

type CardRanking map[byte]int
type HandAndStrength struct {
	strength HandStrength
	hand     string
}

type Hand struct {
	hand string
	bid  int
}

type HandStrength int

const (
	FiveOfAKind  HandStrength = 6
	FourOfAKind  HandStrength = 5
	FullHouse    HandStrength = 4
	ThreeOfAKind HandStrength = 3
	TwoPair      HandStrength = 2
	OnePair      HandStrength = 1
	HighCard     HandStrength = 0
)

func GetStrengthOfHand(hand string) (strength HandStrength) {
	cardMap := map[byte]int{}
	for len(hand) > 0 {
		card := hand[0]
		hand = hand[1:]
		cardMap[card]++
	}

	countMap := map[int][]byte{}
	for card, count := range cardMap {
		countMap[count] = append(countMap[count], card)
	}

	if len(countMap[5]) > 0 {
		return FiveOfAKind
	}
	if len(countMap[4]) > 0 {
		return FourOfAKind
	}
	if len(countMap[3]) > 0 && len(countMap[2]) > 0 {
		return FullHouse
	}
	if len(countMap[3]) > 0 {
		return ThreeOfAKind
	}
	if len(countMap[2]) > 1 {
		return TwoPair
	}
	if len(countMap[2]) > 0 {
		return OnePair
	}

	return HighCard
}

func ReplaceJoker(hand string, cardRanking CardRanking) []string {
	var hands []string
	for card := range cardRanking {
		if card != 'J' {
			hands = append(hands, strings.ReplaceAll(hand, "J", string(card)))
		}
	}
	return hands
}

func GetStrengthOfHandWithJokers(
  hand string, 
  cardRanking CardRanking, 
  cache *map[string]HandStrength,
) HandStrength {
	var jokerCount = 0
	for _, card := range hand {
		if card == 'J' {
			jokerCount++
		}
	}
	if jokerCount == 0 {
		strength := GetStrengthOfHand(hand)
		return HandStrength(jokerCount + int(strength))
	}
	if jokerCount == 5 {
		return FiveOfAKind
	}
  if cachedStrength, ok := (*cache)[hand]; ok {
    return cachedStrength
  }

	possibleHands := ReplaceJoker(hand, cardRanking)
	maxStrength := slices.MaxFunc(possibleHands, func(a, b string) int {
		if GetStrengthOfHand(a) < GetStrengthOfHand(b) {
			return -1
		}
		return 1
	})

  strength :=	GetStrengthOfHand(maxStrength)
  (*cache)[hand] = strength
  return strength
}
func cardComp(a, b string, cardRanking CardRanking) bool {
	for len(a) > 0 {
		cardRankA := cardRanking[a[0]]
		cardRankB := cardRanking[b[0]]

		if cardRankA == cardRankB {
			a = a[1:]
			b = b[1:]
		} else {
			return cardRankA < cardRankB
		}
	}
	return false
}

func handComp(a, b HandAndStrength, cardRanking CardRanking) bool {
	if a.strength == b.strength {
		return cardComp(a.hand, b.hand, cardRanking)
	}
	return a.strength < b.strength
}

func sumOfBids(hands []Hand) int {
	sum := 0
	for i, hand := range hands {
		sum += hand.bid * (i + 1)
	}
	return sum
}

func parseInput(input []string) []Hand {
	var hands = []Hand{}
	for _, line := range input {
		hand, bid := ParseLine(line)
		hands = append(hands, Hand{hand, bid})
	}
	return hands
}
