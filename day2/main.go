package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/jesse-rb/aoc2025/util"
)

func main() {
	lines := util.GetLinesFromStdin()
	part1(lines)
	part2(lines)
}

func part2(lines []string) {
	log.Println("part2")
}

func part1(lines []string) {
	log.Println("part1")

	l := lines[0]

	ranges := strings.Split(l, ",")
	chWorkers := make(chan int, len(ranges))

	// Iterate over ranges
	for _, r := range ranges {
		start_end := strings.Split(r, "-")
		start := start_end[0]
		end := start_end[1]

		s, err := strconv.Atoi(start)
		util.Must(err)

		e, err := strconv.Atoi(end)
		util.Must(err)

		go part1_worker(s, e, chWorkers)
	}

	// Wait for all workers to finish before continuing
	totalInvalidIDs := 0
	for range len(ranges) {
		invalidIDs := <-chWorkers
		totalInvalidIDs += invalidIDs
	}

	log.Printf("%d", totalInvalidIDs)
}

func part1_worker(s int, e int, chWorkers chan int) {
	invalidIDs := 0

	// Iterate over range of IDs to find a ids which consists ONLY of a sequence twice
	for i := s; i <= e; i++ {
		success := true

		curr := strconv.Itoa(i)
		// Must be of even length
		if (len(curr) % 2) != 0 {
			success = false
			continue
		}

		// Byte at index j and index j+len/2 must be equal
		for j := range len(curr) / 2 {
			a := curr[j]
			b := curr[j+len(curr)/2]

			if a != b {
				success = false
				break
			}
		}

		if success {
			// Now we know it must be an invalid ID
			log.Printf("curr: '%s'\n", curr)
			invalidIDs += i
		}
	}

	chWorkers <- invalidIDs
}
