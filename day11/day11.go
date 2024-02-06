package main

import (
	"advent_of_code_2023/utils"
	"fmt"
	"slices"
)

func main() {
	filePath := utils.GetFilePathFromArgs()
	input := utils.ReadFileToArray(filePath, "\n")
	input = input[:len(input)-1]

	Part1(input)
	Part2(input)
}

func Part1(input []string) {
	emptyRows, emptyCols := getEmptyIndexes(input)
	planetLocations := getPlanetLocations(input)
	expansionScale := 2
	total := getSumOfShortestPath(planetLocations, expansionScale, emptyRows, emptyCols)
	fmt.Println("Part1", total)
}

func Part2(input []string) {
	emptyRows, emptyCols := getEmptyIndexes(input)
	planetLocations := getPlanetLocations(input)
	expansionScale := 1000000
	total := getSumOfShortestPath(planetLocations, expansionScale, emptyRows, emptyCols)
	fmt.Println("Part2", total)
}

type Point struct {
	x int
	y int
}


func getEmptyIndexes(input []string) ([]int, []int) {
	emptyRows := getRowExpansionIndexes(input)
	emptyColmns := getColExpansionIndexes(input)
	return emptyRows, emptyColmns
}

func getRowExpansionIndexes(input []string) []int {
	rows := len(input)
	cols := len(input[0])

	emptyRowsIndexes := []int{}
	for j := 0; j < cols; j++ {
		row := ""
		for i := 0; i < rows; i++ {
			row += string(input[i][j])
		}

		if utils.Every('.', []rune(row)) {
			emptyRowsIndexes = append(emptyRowsIndexes, j)
		}
	}
	return emptyRowsIndexes
}

func getColExpansionIndexes(input []string) []int {
	emptyColsIndexes := []int{}
	for i, row := range input {
		if utils.Every('.', []rune(row)) {
			emptyColsIndexes = append(emptyColsIndexes, i)
		}
	}
	return emptyColsIndexes
}

func getPlanetLocations(input []string) []Point {
	points := []Point{}
	for y, row := range input {
		for x, col := range row {
			if col == '#' {
				points = append(points, Point{x, y})
			}
		}
	}
	return points
}

func getSumOfShortestPath(
	planetLocations []Point,
	expansionScale int,
	emptyRows []int,
	emptyColumns []int,
) int {
	total := 0
	for index, p1 := range planetLocations {
		for _, p2 := range planetLocations[:index] {

			// rows
			for i := min(p1.x, p2.x); i < max(p1.x, p2.x); i++ {
				total += getExpansion(emptyRows, i, expansionScale)
			}
			// cols
			for i := min(p1.y, p2.y); i < max(p1.y, p2.y); i++ {
				total += getExpansion(emptyColumns, i, expansionScale)
			}
		}
	}
	return total
}

func getExpansion(emptyIndexes []int, index int, expansionScale int) int {
	if slices.Contains(emptyIndexes, index) {
		return expansionScale
	}
	return 1
}
