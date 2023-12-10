package main

import (
	"advent_of_code_2023/utils"
	"fmt"
	"math"
	"strings"
	"time"
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

	Part1Constant(input)
	Part1(input)

	start2 := time.Now()
	Part2(input)
	elapsed2 := time.Since(start2)
	fmt.Println("Part 2 took", elapsed2)

	start22 := time.Now()
	Part2Constant(input)
	elapsed22 := time.Since(start22)
	fmt.Println("Part 2 Const took", elapsed22)
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
		plays := 0
		time := utils.DangerouslyParseInt(times[i])
		record := utils.DangerouslyParseInt(distances[i])
		for j := 0; j < time; j++ {
			var timeLeft = time - j
			var distance = timeLeft * j
			if distance > record {
				plays += 1
			}
		}

		fmt.Println(time, record, "plays", plays)
		totalPlays = append(totalPlays, plays)
	}

	fmt.Println("Part 1: ", Product(totalPlays))
}

func Part1Constant(input []string) {
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
	distance := strings.Join(strings.Fields(input[1])[1:], "")

	time := utils.DangerouslyParseInt(t)
	record := utils.DangerouslyParseInt(distance)
	plays := 0
	for j := 0; j < time; j++ {
		var timeLeft = time - j
		var distance = timeLeft * j
		if distance > record {
			plays += 1
		}
	}

	fmt.Println(time, record, "plays", plays)

	fmt.Println("Part 2: ", plays)
}

func Part2Constant(input []string) {
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

	fmt.Println(time, record, "plays", plays)
	fmt.Println("Part 2: ", plays)
}
