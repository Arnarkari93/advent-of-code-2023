package main

import (
	"advent_of_code_2023/utils"
	"fmt"
	"slices"
	"sort"
	"strings"
)

func main() {
	filePath := utils.GetFilePathFromArgs()
	input := utils.ReadFileToArray(filePath, "\n")
	input = input[:len(input)-1]
	Part1(input)
}

func ParseLine(line string) (string, int) {
	lineSplit := strings.Fields(line)
	hand, bid := lineSplit[0], utils.DangerouslyParseInt(lineSplit[1])
	return hand, bid
}

type HandStrength int

const (
	FiveOfAKind  HandStrength = 0
	FourOfAKind  HandStrength = 1
	FullHouse    HandStrength = 2
	ThreeOfAKind HandStrength = 3
	TwoPair      HandStrength = 4
	OnePair      HandStrength = 5
	HighCard     HandStrength = 6
	None         HandStrength = 7
)

func GetHandStrength(hand string) (strength HandStrength) {
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

type Hand struct {
	hand string
	bid  int
}

var AllCards = []byte{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}

func handComp(a, b string) bool {
	strengthA := GetHandStrength(a)
	strengthB := GetHandStrength(b)
	if strengthA == strengthB {
		for len(a) > 0 {
			indexA := slices.Index(AllCards, a[0])
			indexB := slices.Index(AllCards, b[0])

			if indexA == indexB {
				a = a[1:]
				b = b[1:]
			} else {
				return indexA < indexB
			}
		}
	}
	return strengthA < strengthB

}

func Part1(input []string) {
	var hands = []Hand{}
	for _, line := range input {
		hand, bid := ParseLine(line)
		hands = append(hands, Hand{hand, bid})
	}

	sort.Slice(hands, func(a, b int) bool {
    return handComp(hands[a].hand, hands[b].hand)
	})

	sum := 0
	for i, hand := range hands {
		sum += hand.bid * (len(hands) - i)
	}

	// fmt.Println(hands)
	fmt.Println(sum)
}
