package main

import (
	"log"
	"strconv"

	"github.com/jesse-rb/aoc2025/util"
)

func main() {
	lines := util.GetLinesFromStdin()
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	debug := false
	// debug := true
	log.Println("PART ONE")

	sum := 0
	for _, bank := range lines {
		largestPair := 0

		left := '0'
		// right := '0'

		for i, current := range bank {
			if current > left && i < len(bank)-1 {
				left = current
				// right = rune(bank[i+1])
				continue
			}
			// jolt := int(joltRune - '0')
			pair, err := strconv.Atoi(string([]rune{left, current}))
			if err != nil {
				panic("Failed to convert string to int")
			}

			if pair > largestPair {
				// right = current
				largestPair = pair
			}

		}

		if debug {
			log.Printf("largest pair: %d", largestPair)
		}

		sum += largestPair
	}

	log.Printf("Sum of highest bank pairs: %d", sum)
}

func part2(lines []string) {
	// debug := false;
	// debug := true
	log.Println("PART TWO")
}
