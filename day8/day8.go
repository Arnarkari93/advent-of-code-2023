package main

import (
	"advent_of_code_2023/utils"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// fmt.Println("Hello, World!")
	filePath := utils.GetFilePathFromArgs()
	input := utils.ReadFileToArray(filePath, "\n")
	input = input[:len(input)-1]

	Part1(input)
	Part2(input)
}

type Node struct {
	left  string
	right string
}

func getLocation(node Node, instruction byte) string {
	if instruction == 'L' {
		return node.left
	}
	return node.right
}

func Part1(input []string) {
	instructions := input[0]
	networkMap, _ := getNetworkMap(input[2:])

	var location = "AAA"
	var steps = 1
	for {
		instruction := instructions[(steps-1)%len(instructions)]
		location = getLocation(networkMap[location], instruction)

		if location == "ZZZ" {
			break
		}
		steps++
	}
	fmt.Println(location, steps)
}

type Path struct {
	startingPoint string
	location      string
	stepsToEnd    int
}

func Part2(input []string) {
	instructions := input[0]
	networkMap, startingLocations := getNetworkMap(input[2:])

	var path = []Path{}
	for _, location := range startingLocations {
		path = append(path, Path{location, location, 0})
	}

	steps := 1
	for {
		instruction := instructions[(steps-1)%len(instructions)]
		var allPathsStepsCalculated = true
		for i, l := range path {
			if path[i].stepsToEnd != 0 {
				continue
			}

			path[i].location = getLocation(networkMap[l.location], instruction)

			if endsWithZ := strings.HasSuffix(path[i].location, "Z"); !endsWithZ {
				allPathsStepsCalculated = false
			} else {
				path[i].stepsToEnd = steps
			}
		}
		if allPathsStepsCalculated {
			break
		}
		steps++
	}

	numbers := []int{}
	for _, l := range path {
		numbers = append(numbers, l.stepsToEnd)
	}

	fmt.Println("Steps", steps)
  fmt.Println("Path", path)
	fmt.Println("lcm", lcmOfSlice(numbers))
}

// Function to calculate the Greatest Common Divisor
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Function to calculate the Least Common Multiple of two numbers
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// Function to calculate the LCM of a slice of numbers
func lcmOfSlice(numbers []int) int {
	result := numbers[0]
	for _, number := range numbers[1:] {
		result = lcm(result, number)
	}
	return result
}

func getNetworkMap(navigations []string) (map[string]Node, []string) {
	var theMap = map[string]Node{}
	var startingLocations = []string{}
	for _, navigation := range navigations {
		value, left, right := parseNavigation(navigation)
		theMap[value] = Node{left, right}
		if strings.HasSuffix(value, "A") {
			startingLocations = append(startingLocations, value)
		}
	}

	return theMap, startingLocations
}

var pattern = regexp.MustCompile(`(\w+)\s*=\s*\((\w+),\s*(\w+)\)`)

func parseNavigation(navigation string) (value, left, right string) {
	matches := pattern.FindStringSubmatch(navigation)
	return matches[1], matches[2], matches[3]
}
