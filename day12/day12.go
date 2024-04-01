package main

import (
	"advent_of_code_2023/utils"
	"fmt"
	"strings"
)

func main() {
	filePath := utils.GetFilePathFromArgs()
	input := utils.ReadFileToArray(filePath, "\n")
	input = input[:len(input)-1]

	Part2(input)
}

func ParseLine(line string) (string, []int) {
	fields := strings.Fields(line)
	springs, groupsString := fields[0], fields[1]

	var groups []int
	for _, num := range strings.Split(groupsString, ",") {
		groups = append(groups, utils.DangerouslyParseInt(num))
	}

	return springs, groups
}

func Part1(input []string) {
	sum := 0
	for _, line := range input {
		springs, groups := ParseLine(line)
		m := make(map[string]int)
		arrangements := SpringsArrangements(springs, groups, m)
		sum += arrangements
	}

	fmt.Println("Total", sum)
}

func Part2(input []string) {
	sum := 0
	for i, line := range input {
		springs, groups := ParseLine(line)
		m := make(map[string]int)
    unfoldedSprings, unfoldedGroups := Unfold(springs, groups)
		arrangements := SpringsArrangements(unfoldedSprings, unfoldedGroups, m)
    fmt.Println(i, arrangements)
		sum += arrangements
	}
	fmt.Println("Total", sum)
}

func Unfold(springs string, groups []int) (string, []int) {
  springsRepeted := strings.Repeat(springs+"?", 5)
  unfoldedSprings := springsRepeted[:len(springsRepeted)-1]
  unfoldedGroups := []int{}
  for i := 0; i < 5; i++ {
    unfoldedGroups = append(unfoldedGroups, groups...) 
  }
  return unfoldedSprings, unfoldedGroups
}

func SpringsArrangements(springs string, groups []int, memo map[string]int) int {
	if len(groups) == 0 {
		if !strings.Contains(springs, "#") {
			return 1
		}
		return 0
	}

	// not enough springs to satisfy the groups
	if len(springs) < (Sum(groups) + len(groups) - 1) {
		return 0
	}

	key := springs + fmt.Sprintf("%v", groups)
	if _, found := memo[key]; found {
		return memo[key]
	}

	// if current spring is operational we skip to the next one
	if springs[0:1] == "." {
		return SpringsArrangements(springs[1:], groups, memo)
	}

	total := 0
	currentGroup := groups[0]

	if !strings.Contains(springs[0:currentGroup], ".") &&
		(len(springs) == currentGroup || springs[currentGroup:currentGroup+1] != "#") {
		next := min(len(springs), currentGroup+1)
		total += SpringsArrangements(springs[next:], groups[1:], memo)
	}

	if springs[0:1] != "#" {
		total += SpringsArrangements(springs[1:], groups, memo)
	}

	memo[key] = total
	return total
}

func Sum(groups []int) int {
	sum := 0
	for _, num := range groups {
		sum += num
	}
	return sum
}
