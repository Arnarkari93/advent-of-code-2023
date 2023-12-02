package main

import (
	"advent_of_code_2023/utils"
	"fmt"
	"strings"
)

type cubes = map[string]int

func main() {
	filePath := utils.GetFilePathFromArgs()
	input := utils.ReadFileToArray(filePath)
	utils.PrintInput(input)

	part1(input, cubes{"red": 12, "green": 13, "blue": 14})
	part2(input)
}

func getGameNumber(gameLabel string) int {
	return utils.DangerouslyParseInt(gameLabel[5:])
}

func parseDraw(draw string) map[string]int {
	drawnCubes := cubes{"red": 0, "green": 0, "blue": 0}

	cubesInDraw := strings.Split(draw, ",")
	for _, cubes := range cubesInDraw {
		count, color, successfulCut := strings.Cut(strings.TrimSpace(cubes), " ")
		if !successfulCut {
			panic("Failed to parse cube, " + draw)
		}

		drawnCubes[color] = utils.DangerouslyParseInt(count)
	}
	return drawnCubes
}

func isDrawPossible(draw string, maxCubes cubes) bool {
	drawnCubes := parseDraw(draw)

	for color, count := range drawnCubes {
		if count > maxCubes[color] {
			return false
		}
	}

	return true
}

func parseGameLine(gameLine string) (int, string) {
	gameLabel, game, successfulCut := strings.Cut(gameLine, ":")
	if !successfulCut {
		panic("Failed to parse game line, " + gameLine)
	}
	gameNumber := getGameNumber(gameLabel)
	return gameNumber, game
}

func part1(input []string, maxCubes cubes) {
	fmt.Println("Part 1")
	fmt.Println("------")
	fmt.Println("Red cubes: ", maxCubes["red"])
	fmt.Println("Green cubes: ", maxCubes["green"])
	fmt.Println("Blue cubes: ", maxCubes["blue"])

	utils.PrintInput(input)

	var sum = 0
	for _, gameLine := range input {
		if len(gameLine) == 0 {
			continue
		}

		gameNumber, game := parseGameLine(gameLine)

		draws := strings.Split(game, ";")

		validGame := true
		for _, draw := range draws {
			validGame = isDrawPossible(draw, maxCubes)
			if !validGame {
				fmt.Println(gameLine)
				fmt.Println("Impssible draw", draw)
				break
			}
		}

		if validGame {
			sum += gameNumber
		}
	}

	fmt.Println("Part 1 sum:", sum)
}

func part2(input []string) {
	fmt.Println("Part 2")
	fmt.Println("------")

	var sum = 0
	for _, gameLine := range input {
		if len(gameLine) == 0 {
			continue
		}

		_, game := parseGameLine(gameLine)
		draws := strings.Split(game, ";")

		minumumCubes := cubes{"red": 0, "green": 0, "blue": 0}
		for _, draw := range draws {
			cubesInDraw := parseDraw(draw)
			for color, count := range cubesInDraw {
				if count > minumumCubes[color] {
					minumumCubes[color] = count
				}
			}
		}

		sum += minumumCubes["red"] * minumumCubes["green"] * minumumCubes["blue"]
	}

	fmt.Println("Part 2 sum:", sum)
}
