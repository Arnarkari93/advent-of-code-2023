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

func Every[T comparable](value T, list []T) bool {
  for _, item := range list {
    if item != value {
      return false
    }
  }
  return true
}

func Filter[T comparable](input []T, value T) []T {
  filtered := []T{}
  for _, s := range input {
    if s != value {
      filtered = append(filtered, s)
    }
  }
  return filtered
}

func FilterMany[T comparable](input []T, values []T) []T {
  // Create a map to store values to be filtered out
    filterMap := make(map[T]bool)

    // Add values to be filtered out to the map
    for _, value := range values {
        filterMap[value] = true
    }

    // Create a slice to store the filtered values
    result := make([]T, 0, len(input))

    // Iterate through the array and append non-filtered values to the result slice
    for _, value := range input {
        if !filterMap[value] {
            result = append(result, value)
        }
    }

    return result
}
