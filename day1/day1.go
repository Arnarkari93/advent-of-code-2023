package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
  "advent_of_code_2023/utils"
)

func main() {

  filePath := utils.GetFilePathFromArgs()
	input := utils.ReadFileToArray(filePath)
  utils.PrintInput(input)

	if strings.Contains(os.Args[1], "part1") {
		part1(input)
		return
	}

	var allowedNumbers = []string{
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}
	if strings.Contains(os.Args[1], "part2") {
		part2(input, allowedNumbers)
		return
	}

	part1(input)
	part2(input, allowedNumbers)
}

func findFirstNumberInPuzzle(puzzle string, matchWords []string) int {
	for i, char := range puzzle {

		if unicode.IsDigit(rune(char)) {
			return int(char - '0')
		}

		for j, word := range matchWords {
			if len(puzzle[i:]) >= len(word) && strings.EqualFold(puzzle[i:i+len(word)], word) {
				return j + 1
			}
		}
	}
	return -1
}

func findLastNumberInPuzzle(str string, matchWords []string) int {
	for i := len(str) - 1; i >= 0; i-- {

		if unicode.IsDigit(rune(str[i])) {
			return int(str[i] - '0')
		}

		for number, word := range matchWords {
			if len(str)-i >= len(word) &&
				strings.EqualFold(str[i:i+len(word)], word) {
				return number + 1
			}
		}
	}
	return -1
}

func part1(input []string) {
	var sum = 0
	for _, puzzle := range input {
		if len(puzzle) == 0 {
			continue
		}

		firstNumber := findFirstNumberInPuzzle(puzzle, []string{})
		lastNumber := findLastNumberInPuzzle(puzzle, []string{})

		number, err := strconv.Atoi(string(firstNumber) + string(lastNumber))
		if err != nil {
			fmt.Println("Error for puzzle", puzzle)
			fmt.Println(err)
			continue
		}

		fmt.Println(number)
		sum += number
	}
	fmt.Println(sum)
}

func part2(intput []string, allowedNumbers []string) {
	var sum = 0
	for _, puzzle := range intput {
		if len(puzzle) == 0 {
			continue
		}

		firstNumber := findFirstNumberInPuzzle(puzzle, allowedNumbers)
		lastNumber := findLastNumberInPuzzle(puzzle, allowedNumbers)

		twoDigitNumber, err := strconv.Atoi(strconv.Itoa(firstNumber) + strconv.Itoa(lastNumber))
		if err != nil {
			fmt.Println("Error for puzzle", puzzle)
			fmt.Println(err)
			continue
		}

		fmt.Print(puzzle, ": ")
		fmt.Println(twoDigitNumber)
		sum += twoDigitNumber
	}
	fmt.Println(sum)
}
