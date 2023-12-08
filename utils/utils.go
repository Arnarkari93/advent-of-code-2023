package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadFile(filePath string) []byte {
	data, err := os.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	return data
}

func ReadFileToArray(filepath string, split string) []string {
	data := ReadFile(filepath)
  if split == "" {
    return strings.Split(string(data), "\n")
  }
	return strings.Split(string(data), split)
}

func PrintInput(input []string) {
	for _, puzzle := range input {
		fmt.Println(puzzle)
	}
}

func GetFilePathFromArgs() string {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run day.go <input.txt>")
		panic("No input file provided")
	}

	return os.Args[1]
}

func DangerouslyParseInt(input string) int {
	number, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return number
}

func StringNumbersToIntArray(input string) []int {
  input = strings.TrimSpace(input)
  numbers := []int{}
  for _, s := range strings.Split(input, " ") {
    number, _ := strconv.Atoi(s)
    numbers = append(numbers, number)
  }
  return numbers
}
