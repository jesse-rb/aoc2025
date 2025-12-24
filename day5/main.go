package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/jesse-rb/aoc2025/util"
)

func main() {
	lines := util.GetLinesFromStdin()
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	freshRanges := make([][2]int, 0)

	nextProcessing := 0

	numFresh := 0

	for _, l := range lines {
		if l == "" {
			nextProcessing++
			continue
		}

		if nextProcessing < 1 {
			// Build fresh ID ranges
			ranges := strings.Split(l, "-")
			from, _ := strconv.Atoi(ranges[0])
			to, _ := strconv.Atoi(ranges[1])
			freshRanges = append(freshRanges, [2]int{from, to})
		} else {
			// Process IDs
			id, _ := strconv.Atoi(l)

			// Attempt to find the id in the fresh ranges
			for _, r := range freshRanges {
				if id >= r[0] && id <= r[1] {
					numFresh++
					break
				}
			}
		}
	}

	log.Printf("Num fresh: %d\n", numFresh)
}

// Collect ranges into tuple slices
// Sort by start in Asc order
// Build non-overlapping ranges
func part2(lines []string) {
	log.Println("part TWO")

	numFresh := 0
	ranges := make([][2]int, 0)

	for _, l := range lines {
		if l == "" {
			break
		}

		// Build fresh ID ranges
		r := strings.Split(l, "-")
		from, _ := strconv.Atoi(r[0])
		to, _ := strconv.Atoi(r[1])

		ranges = append(ranges, [2]int{from, to})
	}

	// Sort ranges
	slices.SortFunc(ranges, func(a [2]int, b [2]int) int {
		return a[0] - b[0]
	})

	// Build non-overlapping ranges
	// nonOverlappingRanges := make([][2]int, 0)
	currentRange := [2]int{}
	for i, r := range ranges {
		if i == 0 {
			currentRange = r
		}

		if i < len(ranges)-1 {
			rNext := ranges[i+1]
			// If overlaps with next range
			if currentRange[1] > rNext[0] {
				// nonOverlappingRanges = append(nonOverlappingRanges, [2]int{currentRange[0], rNext[0]})
				numFresh += (currentRange[1] - 1) - currentRange[0] + 1
				currentRange = [2]int{currentRange[1], rNext[1]}
			} else {
				// nonOverlappingRanges = append(nonOverlappingRanges, currentRange)
				numFresh += currentRange[1] - currentRange[0] + 1
				currentRange = rNext
			}
		} else {
			numFresh += currentRange[1] - currentRange[0] + 1
			// nonOverlappingRanges = append(nonOverlappingRanges, currentRange)
		}
	}

	// SUM
	// for _, r := range nonOverlappingRanges {
	// 	// log.Printf("non-overlapping range: %d - %d", r[0], r[1])
	// 	numFresh += r[1] - r[0] + 1
	// }

	log.Printf("Num fresh: %d\n", numFresh)
}

func part2MapKey(from int, to int) string {
	return fmt.Sprintf("%d-%d", from, to)
}
