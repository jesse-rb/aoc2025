package main

import (
	"fmt"
	"log"

	"github.com/jesse-rb/aoc2025/util"
)

func main() {
	lines := util.GetLinesFromStdin()
	part1(lines)
	part2(lines)
}

func checkCoord(x int, y int, lines []string, cache map[string]byte) bool {
	cacheKey := fmt.Sprintf("%d,%d", x, y)

	if cached, exists := cache[cacheKey]; exists {
		return cached == '@'
	}

	if y < 0 || y > len(lines)-1 {
		return false
	}
	if x < 0 || x > len(lines[y])-1 {
		return false
	}

	cache[cacheKey] = lines[y][x]

	return lines[y][x] == '@'
}

func part1(lines []string) {
	log.Println("Part ONE")

	// debug := true
	// debug := true

	solution := 0

	cache := make(map[string]byte, 0)
	for y := range lines {
		for x := range lines[y] {
			if lines[y][x] != '@' {
				continue
			}

			adjecentCoords := [8][2]int{
				{x, y - 1},
				{x + 1, y - 1},
				{x + 1, y},
				{x + 1, y + 1},
				{x, y + 1},
				{x - 1, y + 1},
				{x - 1, y},
				{x - 1, y - 1},
			}

			numPaperRolls := 0
			for _, coords := range adjecentCoords {
				if checkCoord(coords[0], coords[1], lines, cache) {
					numPaperRolls++
				}
			}

			if numPaperRolls < 4 {
				solution++
			}
		}
	}

	log.Printf("total accessible paper rolls: %d", solution)
}

func part2(lines []string) {
	log.Println("Part TWO")

	// debug := true
	// debug := true

	solution := 0

	cache := make(map[string]byte, 0)

	isInitial := true
	canRemove := 0

	for canRemove > 0 || isInitial {
		isInitial = false
		canRemove = 0

		for y := range lines {
			for x := range lines[y] {
				if lines[y][x] != '@' {
					continue
				}

				adjecentCoords := [8][2]int{
					{x, y - 1},
					{x + 1, y - 1},
					{x + 1, y},
					{x + 1, y + 1},
					{x, y + 1},
					{x - 1, y + 1},
					{x - 1, y},
					{x - 1, y - 1},
				}

				numPaperRolls := 0
				for _, coords := range adjecentCoords {
					if checkCoord(coords[0], coords[1], lines, cache) {
						numPaperRolls++
					}
				}

				if numPaperRolls < 4 {
					lineRuneSlice := []rune(lines[y])
					lineRuneSlice[x] = '.'
					lines[y] = string(lineRuneSlice)

					cacheKey := fmt.Sprintf("%d,%d", x, y)
					cache[cacheKey] = lines[y][x]
					canRemove++
				}
			}
		}

		solution += canRemove
	}

	log.Printf("Can remove %d paper rolls.", solution)
}
