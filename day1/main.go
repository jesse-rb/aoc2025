package main

import (
	"log"
	"strconv"

	"github.com/jesse-rb/aoc2025/util"
)

func main() {
	log.Println("Hello world")

	current := 50

	countCurrentLandOnZero := 0

	lines := util.GetLinesFromStdin()
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
