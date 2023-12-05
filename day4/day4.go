package main

import (
	"advent_of_code_2023/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	filePath := utils.GetFilePathFromArgs()
	input := utils.ReadFileToArray(filePath)

	// part1(input)
	start2 := time.Now()
	part2(input)
	elapsed2 := time.Since(start2)
	fmt.Println("Part 2 took", elapsed2)
}

func makeSortedList(str string) []int {
	trimmed := strings.Trim(str, " ")
	split := strings.Split(trimmed, " ")

	numbers := []int{}
	for _, s := range split {
		if number, err := strconv.Atoi(s); err == nil {
			numbers = append(numbers, number)
		}
	}

	slices.Sort(numbers)
	return numbers
}

func calculatePoints(points int) int {
	if points == 0 {
		return 1
	}
	return points * 2
}

func pop(slice *[]int) {
	*slice = (*slice)[1:]
}

func peek(slice []int) int {
	return slice[0]
}

func parseCard(game string) (string, string, string, bool) {
	label, numbers, succesfulCut := strings.Cut(game, ":")
	if !succesfulCut {
		return "", "", "", false
	}
	winingNumbers, ourNumbers, succesfulCut := strings.Cut(numbers, "|")
	if !succesfulCut {
		return "", "", "", false
	}
	return label, winingNumbers, ourNumbers, true
}


func part1(input []string) {
	var sum = 0
	for _, line := range input {
		label, winingNumbers, ourNumbers, succesfulCut := parseCard(line)
		if !succesfulCut {
			continue
		}

		winningList := makeSortedList(winingNumbers)
		ourNumbersList := makeSortedList(ourNumbers)

		var points = 0
		for len(winningList) != 0 && len(ourNumbersList) != 0 {
			winning := peek(winningList)
			our := peek(ourNumbersList)

			if winning == our {
				pop(&winningList)
				pop(&ourNumbersList)
				points = calculatePoints(points)
			} else if winning > our {
				pop(&ourNumbersList)
			} else {
				pop(&winningList)
			}
		}
		sum += points
		fmt.Println(label, points)
	}
	fmt.Println("Sum:", sum)
}

func getScratchCardCount(cardNumber int, scratchcards map[int]int) int {
	var count = scratchcards[cardNumber]
	if count == 0 {
		return 1
	}
	return count
}

func getCardNumber(gameLabel string) int {
	_, number, succesfulCut := strings.Cut(gameLabel, " ")
	if !succesfulCut {
		return -1
	}
	return utils.DangerouslyParseInt(strings.Trim(number, " "))
}

func getNumberOfMachingNumbers(winningNumbers []int, ourNumbers []int) int {
	var matchingNumbers = 0
	for len(winningNumbers) != 0 && len(ourNumbers) != 0 {
		winning := peek(winningNumbers)
		our := peek(ourNumbers)

		if winning == our {
			matchingNumbers += 1
			pop(&winningNumbers)
			pop(&ourNumbers)
		} else if winning > our {
			pop(&ourNumbers)
		} else {
			pop(&winningNumbers)
		}
	}
	return matchingNumbers
}

func part2(input []string) {
	var sum = 0
	var scratchcards = map[int]int{}
	for _, line := range input {
		label, winingNumbers, ourNumbers, succesfulCut := parseCard(line)
		if !succesfulCut {
			continue
		}
		// set orignal card holdings to 1 before checking for matches
		scratchcards[getCardNumber(label)] += 1

		winningList := makeSortedList(winingNumbers)
		ourNumbersList := makeSortedList(ourNumbers)
    matchingNumbers := getNumberOfMachingNumbers(winningList, ourNumbersList)

    gameNumber := getCardNumber(label)
		// fmt.Println(label, ":", matchingNumbers, matchingNumbers + gameNumber)
		for i := gameNumber + 1; i <= matchingNumbers+gameNumber; i++ {
			scratchcards[i] = scratchcards[i] + scratchcards[gameNumber]
			// fmt.Println("adding ", scratchcards[gameNumber], " to ", i, "=", scratchcards[i])
		}
		// fmt.Println(scratchcards)
		// add count of scratchcards to sum
		sum += scratchcards[gameNumber]
	}
	// fmt.Println(scratchcards)
	fmt.Println("Sum:", sum)
}
