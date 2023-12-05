package main

import (
	"advent_of_code_2023/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	filePath := utils.GetFilePathFromArgs()
	input := utils.ReadFileToArray(filePath)

	part1(input)
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

func part1(input []string) {
  var sum = 0
	for _, line := range input {
		label, numbers, succesfulCut := strings.Cut(line, ":")
		if !succesfulCut {
			continue
		}

		winingNumbers, ourNumbers, succesfulCut := strings.Cut(numbers, "|")
		if !succesfulCut {
			continue
		}

		winningList := makeSortedList(winingNumbers)
		ourNumbersList := makeSortedList(ourNumbers)

		var points = 0
		for len(winningList) != 0 && len(ourNumbersList) != 0 {
			winning := winningList[0]
			our := ourNumbersList[0]

			if winning == our {
				winningList = winningList[1:]
				if points == 0 {
					points = 1
				} else {
					points = points * 2
				}
			} else if winning > our {
        ourNumbersList = ourNumbersList[1:]
      } else {
        winningList = winningList[1:]
      }
		}
    sum += points
		fmt.Println(label, points)
	}
  fmt.Println("Sum:", sum)
}
