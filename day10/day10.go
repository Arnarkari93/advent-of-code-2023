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

	// Part1(input)
	Part2(input)
}

type Point struct {
	x int
	y int
}

type Tile struct {
	tile              rune
	visited           bool
	connetions        []Direction
	distanceFromStart int
}

type Direction rune

const (
	North Direction = 'N'
	South Direction = 'S'
	East  Direction = 'E'
	West  Direction = 'W'
)

func negateDirection(direction Direction) Direction {
	switch direction {
	case North:
		return South
	case South:
		return North
	case East:
		return West
	case West:
		return East
	default:
		panic("Invalid connection")
	}
}

func getTileConnetions(tile rune) []Direction {
	switch tile {
	case '|':
		return []Direction{North, South}
	case '-':
		return []Direction{East, West}
	case 'F':
		return []Direction{South, East}
	case '7':
		return []Direction{South, West}
	case 'L':
		return []Direction{North, East}
	case 'J':
		return []Direction{North, West}
	default:
		return []Direction{}
	}
}

func move(from Point, direction Direction) Point {
	switch direction {
	case North:
		return Point{from.x, from.y - 1}
	case South:
		return Point{from.x, from.y + 1}
	case East:
		return Point{from.x + 1, from.y}
	case West:
		return Point{from.x - 1, from.y}
	default:
		panic("Invalid direction")
	}
}

func travel(connection Direction, point Point, theMap map[Point]Tile) {
	travelDistance := 0
	for {
		nextPoint := move(point, connection)
		tile, exist := theMap[nextPoint]
		if !exist {
			return // out of bounds
		}

		connectionAt := slices.Index(tile.connetions, negateDirection(connection))
		if connectionAt == -1 {
			return // not a tile connection
		}

		if tile.distanceFromStart > travelDistance || tile.distanceFromStart == 0 {
			tile.distanceFromStart = travelDistance + 1
		}

		tile.visited = true
		theMap[nextPoint] = tile
		point = nextPoint
		travelDistance++

		connection = tile.connetions[1-connectionAt]
	}
}

func MaxInMap(theMap map[Point]Tile) int {
	max := 0
	for _, tile := range theMap {
		if tile.distanceFromStart > max {
			max = tile.distanceFromStart
		}
	}
	return max
}

func Part1(input []string) {
	startingTile, theMap := buildMap(input)

	travel(East, startingTile, theMap)
	travel(South, startingTile, theMap)
	travel(West, startingTile, theMap)
	travel(North, startingTile, theMap)

	max := MaxInMap(theMap)
	fmt.Println("Max", max)
}

func countEgdes(startingPoint Point, theMap map[Point]Tile, direction Direction) int {
	edgeCrossing := 0
	point := move(startingPoint, direction)

	ridingEdge := '.'
	// When we hit a F or L we only count a single edgeCrossing until we hit a
	// matching 7 or L that results in a need to cross over. Other wise we could
	// ride along the edge and we dont count that
	if theMap[point].tile == 'L' || theMap[point].tile == 'F' {
		ridingEdge = theMap[point].tile
	}
	for {
		tile, inBounds := theMap[point]
		if !inBounds {
			break
		}
		if !tile.visited {
			point = move(point, direction)
			ridingEdge = '.'
			continue
		}
		if tile.tile == 'L' || tile.tile == 'F' {
			ridingEdge = tile.tile
		}
		if ridingEdge == 'L' && tile.tile == '7' {
			edgeCrossing++
			ridingEdge = '.'
		}
		if ridingEdge == 'F' && tile.tile == 'J' {
			edgeCrossing++
			ridingEdge = '.'
		}
		if tile.tile == '|' {
			edgeCrossing++
			ridingEdge = '.'
		}
		point = move(point, direction)
	}
	return edgeCrossing
}

func getStartingReplacement(startingPoint Point, theMap map[Point]Tile) []rune {
	var maybeS = []rune{'|', '-', 'F', '7', 'L', 'J'}

	northP := move(startingPoint, North)
	southP := move(startingPoint, South)
	eastP := move(startingPoint, East)
	westP := move(startingPoint, West)

	if theMap[northP].visited && slices.Contains([]rune("F7|"), theMap[northP].tile) {
		maybeS = utils.FilterMany(maybeS, []rune("F7-"))
	}

	if theMap[eastP].visited && slices.Contains([]rune("J7-"), theMap[eastP].tile) {
		maybeS = utils.FilterMany(maybeS, []rune("J7|"))
	}

	if theMap[southP].visited && slices.Contains([]rune("JL|"), theMap[southP].tile) {
		maybeS = utils.FilterMany(maybeS, []rune("JL-"))
	}

	if theMap[westP].visited && slices.Contains([]rune("LF-"), theMap[westP].tile) {
		maybeS = utils.FilterMany(maybeS, []rune("LF|"))
	}

	return maybeS
}

func Part2(input []string) {
	startingTile, theMap := buildMap(input)

	travel(East, startingTile, theMap)
	travel(North, startingTile, theMap)
	travel(South, startingTile, theMap)
	travel(West, startingTile, theMap)

	// Replace the starting tile with the correct tile so we can use
	// ray casting algorithm for potential points going through the starting tile
	sReplacement := getStartingReplacement(startingTile, theMap)

	modifiedStartingTile := theMap[startingTile]
	modifiedStartingTile.tile = sReplacement[0]
	theMap[startingTile] = modifiedStartingTile

	// Use ray casting algorithm to find points inside the map
	pointsInside := []Point{}
	for y, row := range input {
		for x := range row {
			if theMap[Point{x, y}].visited {
				fmt.Print(string(theMap[Point{x, y}].tile))
				continue
			}

			edges := countEgdes(Point{x, y}, theMap, East)
			if edges%2 == 1 {
				fmt.Print("@")
				pointsInside = append(pointsInside, Point{x, y})
			} else {
				fmt.Print(".")
			}

		}
		fmt.Println()
	}

	fmt.Println("pointsInside", len(pointsInside))
}

func buildMap(input []string) (Point, map[Point]Tile) {
	theMap := map[Point]Tile{}
	startingTile := Point{0, 0}
	for y, row := range input {
		for x, tile := range row {
			connections := getTileConnetions(tile)
			visited := false
			if tile == 'S' {
				startingTile = Point{x, y}
				visited = true
			}
			theMap[Point{x, y}] = Tile{tile, visited, connections, 0}
		}
	}
	return startingTile, theMap
}
