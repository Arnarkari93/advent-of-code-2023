package main

import (
	"advent_of_code_2023/utils"
	"cmp"
	"fmt"
	"slices"
	"strings"
)

func main() {
	filePath := utils.GetFilePathFromArgs()

	input := utils.ReadFileToArray(filePath, "\n\n")

	part1(input)
	part2(input)
}

func ReadSeeds(seedInput string) []int {
	seeds := strings.Split(seedInput, ": ")[1]
	return utils.StringNumbersToIntArray(seeds)
}

func MakeAlmanacRanges(almanacRanges string) [][]int {
	ranges := [][]int{}
	for _, almanacRange := range strings.Split(almanacRanges, "\n")[1:] {
		if almanacRange == "" {
			continue
		}
		ranges = append(ranges, utils.StringNumbersToIntArray(almanacRange))
	}
	return ranges
}

func part1(input []string) {
	seeds := ReadSeeds(input[0])

	almanac := input[1:]
	for _, almanacRanges := range almanac {
		ranges := MakeAlmanacRanges(almanacRanges)

		currentSeeds := []int{}
		for _, seed := range seeds {
			foundMapping := false
			for _, r := range ranges {
				dest, source, _range := r[0], r[1], r[2]
				if seed >= source && seed < source+_range {
					currentSeeds = append(currentSeeds, seed-source+dest)
					foundMapping = true
					break
				}
			}
			if !foundMapping {
				currentSeeds = append(currentSeeds, seed)
			}
		}
		seeds = currentSeeds

	}
	fmt.Println("Min seed", slices.Min(seeds))
}

type SeedRange struct {
	start int
	end   int
}

func pop(s *[]SeedRange) SeedRange {
	value := (*s)[0]
	*s = (*s)[1:]
	return value
}

func makeSeedRanges(seedsInput string) []SeedRange {
	seedNumbers := ReadSeeds(seedsInput)
	ranges := []SeedRange{}
	for i := 0; i < len(seedNumbers); i += 2 {
		start, end := seedNumbers[i], seedNumbers[i]+seedNumbers[i+1]
		ranges = append(ranges, SeedRange{start, end})
	}
	return ranges
}

func part2(input []string) {

	seedRanges := makeSeedRanges(input[0])

	almanac := input[1:]
	for _, almanacRanges := range almanac {

		ranges := MakeAlmanacRanges(almanacRanges)

		currentSeedRanges := []SeedRange{}
		for len(seedRanges) > 0 {
			value := pop(&seedRanges)
			start, end := value.start, value.end

			foundMapping := false
			for _, r := range ranges {
				dest, source, _range := r[0], r[1], r[2]
				overlapStart := max(start, source)
				overlapEnd := min(end, source+_range)
				if overlapStart < overlapEnd {
					foundMapping = true
          seedRange := SeedRange{start: overlapStart - source + dest, end: overlapEnd - source + dest}
					currentSeedRanges = append(currentSeedRanges, seedRange)
					if overlapStart > start {
						seedRanges = append(seedRanges, SeedRange{start, overlapStart})
					}
					if end > overlapEnd {
						seedRanges = append(seedRanges, SeedRange{overlapEnd, end})
					}
					break
				}
			}
			if !foundMapping {
				currentSeedRanges = append(currentSeedRanges, SeedRange{start, end})
			}
		}
		seedRanges = currentSeedRanges
	}

	min := slices.MinFunc(seedRanges, func(a, b SeedRange) int {
		return cmp.Compare(a.start, b.start)
	})
	fmt.Println("Min seedRanges", min, "min", min.start)
}
