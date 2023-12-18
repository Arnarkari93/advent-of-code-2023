package main

import (
	"advent_of_code_2023/utils"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	filePath := utils.GetFilePathFromArgs()
	input := utils.ReadFileToArray(filePath, "\n")
	input = input[:len(input)-1]

	Part1(input)
	Part2(input)
}

func Part1(input []string) {
  sum := 0
	for _, line := range input {
		numbers := utils.StringNumbersToIntArray(line)
    next := nextDifference(numbers) + numbers[len(numbers)-1]
    sum += next
		fmt.Println("numbers", next)
	}

  fmt.Println("sum", sum)
}

func Part2(input []string) {
  sum := 0
	for _, line := range input {
		numbers := utils.StringNumbersToIntArray(line)
    prev := numbers[0] - prevDifference(numbers)  
    sum += prev
		fmt.Println("numbers", prev)
	}

  fmt.Println("sum", sum)
}

func nextDifference(numbers []int) int {
	if utils.Every(0, numbers) {
		return 0
	}

	diffs := diffs(numbers)
	next := nextDifference(diffs) + diffs[len(diffs)-1]
  return next 
}

func prevDifference(numbers []int) int {
	if utils.Every(0, numbers) {
		return 0
	}

	diffs := diffs(numbers)
	prev := diffs[0] - prevDifference(diffs) 
  return prev 
}


func diffs(numbers []int) []int {
	diffs := []int{}
	for i := 0; i < len(numbers) - 1; i++ {
		diffs = append(diffs, numbers[i+1]-numbers[i])
	}

	return diffs
}


