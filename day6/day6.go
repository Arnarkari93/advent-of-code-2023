package main

import (
	"advent_of_code_2023/utils"
	"fmt"
	"math"
	"strings"
)

func calcDiscriminant(a, b, c float64) float64 {
	return b*b - 4*a*c
}

func quadraticFormula(a, b, c float64) (float64, float64) {
	d := calcDiscriminant(a, b, c)
	min := (b - math.Sqrt(d)) / (2 * a)
	max := (b + math.Sqrt(d)) / (2 * a)

	return min, max
}

func Product(numbers []int) int {
	value := 1
	for _, i := range numbers {
		value *= i
	}
	return value
}

func main() {
	filePath := utils.GetFilePathFromArgs()
	input := utils.ReadFileToArray(filePath, "\n")

	Part1(input)
	Part2(input)
}

func getDistance(time, record, min int) int {
	timeLeft := time - min
	distance := timeLeft * min
	return distance
}

func Part1(input []string) {
	times := strings.Fields(input[0])[1:]
	distances := strings.Fields(input[1])[1:]

	totalPlays := []int{}
	for i := 0; i < len(times); i++ {
		time := utils.DangerouslyParseInt(times[i])
		record := utils.DangerouslyParseInt(distances[i])

		min, max := quadraticFormula(1, float64(time), float64(record))

		minDistance := getDistance(time, record, int(math.Floor(min)))
		if minDistance == record {
			min += 1
		}
		plays := math.Floor(max) - math.Floor(min)

		totalPlays = append(totalPlays, int(plays))
	}

	fmt.Println("Part 1: ", Product(totalPlays))
}

func Part2(input []string) {
	t := strings.Join(strings.Fields(input[0])[1:], "")
	r := strings.Join(strings.Fields(input[1])[1:], "")

	time := utils.DangerouslyParseInt(t)
	record := utils.DangerouslyParseInt(r)
	min, max := quadraticFormula(1, float64(time), float64(record))

	minDistance := getDistance(time, record, int(math.Floor(min)))
	if minDistance == record {
		min += 1
	}
	plays := math.Floor(max) - math.Floor(min)

	fmt.Println(time, record, "plays", int(plays))
	fmt.Println("Part 2: ", int(plays))
}
