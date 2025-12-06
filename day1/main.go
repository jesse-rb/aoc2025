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

func part2(lines []string) {
	debug := false
	// debug := true

	log.Println("part2")

	current := 50

	countPassedZero := 0

	for _, l := range lines {
		left := string(l[0]) == "L"
		toMoveString := l[1:]
		toMoveInt, err := strconv.Atoi(toMoveString)
		util.Must(err)

		passes := 0
		prev := current
		if left {
			passes = (util.ModInt(0-current, 100) + toMoveInt) / 100
			current = util.ModInt(current-toMoveInt, 100)
		} else {
			passes = (current + toMoveInt) / 100
			current = util.ModInt(current+toMoveInt, 100)
		}

		countPassedZero += passes

		if debug {
			log.Printf("state: moved from(%d) amount(%s) to(%d) with num passed(%d)\n", prev, l, current, passes)
		}
	}

	log.Println(countPassedZero)
}

func part1(lines []string) {
	log.Println("part1")

	current := 50

	countCurrentLandOnZero := 0

	for _, l := range lines {
		left := string(l[0]) == "L"
		toMoveString := l[1:]
		toMoveInt, err := strconv.Atoi(toMoveString)
		util.Must(err)

		if left {
			current = (current - toMoveInt) % 100
		} else {
			current = (current + toMoveInt) % 100
		}

		if current == 0 {
			countCurrentLandOnZero += 1
		}
	}

	log.Println(countCurrentLandOnZero)
}
