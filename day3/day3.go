package main

import (
	"advent_of_code_2023/utils"
	"fmt"
	"unicode"
)

func main() {
	filePath := utils.GetFilePathFromArgs()
	input := utils.ReadFileToArray(filePath)
	part1(input)
	part2(input)
}

type Point struct {
	x, y int
}

func drawEngine(engine map[Point]rune, input []string) {
	for y, line := range input {
		for x, char := range line {
			engine[Point{x, y}] = char
		}
	}
}

func isSymbol(symbol rune) bool {
	return !unicode.IsDigit(symbol) && symbol != '.'
}

func getPointOnMap(point Point, engine map[Point]rune) rune {
	entry, exists := engine[point]
	if exists {
		return entry
	}
	return '.' // default out of bounds as '.'
}

func getNumberOnMap(point Point, engine map[Point]int) int {
	entry, exists := engine[point]
	if exists {
		return entry
	}
	return -1 // no number here
}

func down(point Point) Point {
	return Point{point.x, point.y + 1}
}
func up(point Point) Point {
	return Point{point.x, point.y - 1}
}
func left(point Point) Point {
	return Point{point.x - 1, point.y}
}
func right(point Point) Point {
	return Point{point.x + 1, point.y}
}

func checkIfAdjecentToSymbol(point Point, engine map[Point]rune) bool {
	if isSymbol(getPointOnMap(up(point), engine)) {
		return true
	}
	if isSymbol(getPointOnMap(down(point), engine)) {
		return true
	}
	if isSymbol(getPointOnMap(left(point), engine)) {
		return true
	}
	if isSymbol(getPointOnMap(right(point), engine)) {
		return true
	}
	if isSymbol(getPointOnMap(left(up(point)), engine)) {
		return true
	}
	if isSymbol(getPointOnMap(right(up(point)), engine)) {
		return true
	}
	if isSymbol(getPointOnMap(left(down(point)), engine)) {
		return true
	}
	if isSymbol(getPointOnMap(right(down(point)), engine)) {
		return true
	}
	return false
}

func shouldAddNumberToSum(isAdjecentToSymbol bool, numberQue []rune, char rune, line string, position int) bool {
	var isAtEndOfLine = len(line)-1 == position
	return isAdjecentToSymbol && len(numberQue) > 0 && (!unicode.IsDigit(char) || isAtEndOfLine)
}

func shouldClearNumberQue(numberQue []rune, char rune, line string, position int) bool {
	var isAtEndOfLine = len(line)-1 == position
	return len(numberQue) > 0 && (!unicode.IsDigit(char) || isAtEndOfLine)
}

func part1(input []string) {
	engine := make(map[Point]rune)
	utils.PrintInput(input)
	drawEngine(engine, input)

	var sum = 0
	for y, line := range input {
		var isAdjecentToSymbol = false
		var numberQue = []rune{}

		for x, char := range line {
			p := Point{x, y}

			if unicode.IsDigit(char) {
				numberQue = append(numberQue, char)

				if checkIfAdjecentToSymbol(p, engine) {
					isAdjecentToSymbol = true
				}
			}

			if shouldAddNumberToSum(isAdjecentToSymbol, numberQue, char, line, x) {
				sum += utils.DangerouslyParseInt(string(numberQue))
				numberQue = []rune{}
				isAdjecentToSymbol = false
			}

			if shouldClearNumberQue(numberQue, char, line, x) {
				numberQue = []rune{}
			}
		}
	}

	fmt.Println("Part sum", sum)
}

func getEngineNumberCord(input []string) map[Point]int {
  var engineNumber = make(map[Point]int)

	for y, line := range input {
    var numberQue = []rune{}
    var points = []Point{}

		for x, char := range line {
      if unicode.IsDigit(char) {
        points = append(points, Point{x, y})
        numberQue = append(numberQue, char)
      }

      if len(numberQue) > 0 && (!unicode.IsDigit(char) || len(line)-1 == x) {
        for _, point := range points {
          engineNumber[point] = utils.DangerouslyParseInt(string(numberQue))
        }
        numberQue = []rune{}
        points = []Point{}
      }
		}
	}
  
  return engineNumber
}

func part2(input []string) {
	fmt.Println("Part 2")
	utils.PrintInput(input)

  var pointToNumberMap = getEngineNumberCord(input)

  var sum = 0
  for y, line := range input {
    for x, char := range line {
      if char == '*' {
        var gearPoint = Point{x, y}

        var numbersFound = []int{}
        // look up
        if num, exists := pointToNumberMap[left(up(gearPoint))]; exists {
          numbersFound = append(numbersFound, num)
          
          // check if there is a number to the right of found number
          if _, isSameNumber := pointToNumberMap[up(gearPoint)]; !isSameNumber {
            // there needs to be space between the two numbers
            if num, exists := pointToNumberMap[right(up(gearPoint))]; exists {
              numbersFound = append(numbersFound, num)
            }
          }


        } else if num, exists := pointToNumberMap[up(gearPoint)]; exists {
          numbersFound = append(numbersFound, num)
        } else if num, exists := pointToNumberMap[right(up(gearPoint))]; exists {
          numbersFound = append(numbersFound, num)
        }
        // look left
        if num, exists := pointToNumberMap[left(gearPoint)]; exists {
          numbersFound = append(numbersFound, num)
        }
        // look right
        if num, exists := pointToNumberMap[right(gearPoint)]; exists {
          numbersFound = append(numbersFound, num)
        }
        // look down
        if num, exists := pointToNumberMap[left(down(gearPoint))]; exists {
          numbersFound = append(numbersFound, num)

          // check if there is a number to the right of found number
          if _, isSameNumber := pointToNumberMap[down(gearPoint)]; !isSameNumber {
            // there needs to be space between the two numbers
            if num, exists := pointToNumberMap[right(down(gearPoint))]; exists {
              numbersFound = append(numbersFound, num)
            }
          }
        } else if num, exists := pointToNumberMap[down(gearPoint)]; exists {
          numbersFound = append(numbersFound, num)
        } else if num, exists := pointToNumberMap[right(down(gearPoint))]; exists {
          numbersFound = append(numbersFound, num)
        }

        // found gear
        if len(numbersFound) == 2 {
          sum += numbersFound[0] * numbersFound[1]
        }
      }
    }
  }

  fmt.Println("Part 2 sum", sum)
}
