package main

import (
	"advent_of_code_2023/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func main() {

	filePath := utils.GetFilePathFromArgs()
	input := utils.ReadFileToArray(filePath)
	// utils.PrintInput(input)

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
	start := time.Now()
	part1(input)
	elapsed := time.Since(start)
	fmt.Println("Part 1 took", elapsed)
	start2 := time.Now()
	part2(input, allowedNumbers)
	elapsed2 := time.Since(start2)
	fmt.Println("Part 2 took", elapsed2)
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

		oneNumber := strconv.Itoa(firstNumber) + strconv.Itoa(lastNumber)
		number := utils.DangerouslyParseInt(oneNumber)
		// fmt.Println(number)
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

		// fmt.Print(puzzle, ": ")
		// fmt.Println(twoDigitNumber)
		sum += twoDigitNumber
	}
	fmt.Println(sum)
}
