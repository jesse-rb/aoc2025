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

		for i, current := range bank {
			if current > left && i < len(bank)-1 {
				left = current
				continue
			}
			pair, err := strconv.Atoi(string([]rune{left, current}))
			if err != nil {
				panic("Failed to convert string to int")
			}

			if pair > largestPair {
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
	debug := false
	// debug := true

	log.Println("PART TWO")

	sum := 0
	for _, bank := range lines {
		largestSeq := 0

		left := ""
		if debug {
			log.Printf("---------\nbank: %s", bank)
		}

		for i, current := range bank {
			leftRunes := []rune(left)

			lenLeftRunes := len(leftRunes)
			lenRemaining := len(bank) - i
			spaceRemaining := 12 - lenLeftRunes
			lastIndex := lenLeftRunes - 1

			if i < len(bank)-1 {
				if lastIndex >= 0 && current > leftRunes[lastIndex] && lenRemaining > spaceRemaining {
					offset := 0
					for lastIndex >= 0 && current > leftRunes[lastIndex] && lenRemaining > (spaceRemaining+offset) {
						leftRunes = leftRunes[:lastIndex+1]
						leftRunes[lastIndex] = current
						lastIndex--
						offset++
					}
					left = string(leftRunes)
					if debug {
						log.Printf("leftRunes: %s", string(leftRunes))
					}
					continue
				} else if spaceRemaining > 1 {
					leftRunes = append(leftRunes, current)
					left = string(leftRunes)
					if debug {
						log.Printf("leftRunes: %s", string(leftRunes))
					}
					continue
				}
			}

			seq, err := strconv.Atoi(string(append(leftRunes, current)))
			if err != nil {
				panic("Failed to convert string to int")
			}

			if seq > largestSeq {
				largestSeq = seq
			}

		}

		if debug {
			log.Printf("largest seq: %d", largestSeq)
		}

		sum += largestSeq
	}

	log.Printf("Sum of highest bank sequences: %d", sum)
}
