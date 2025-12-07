package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/jesse-rb/aoc2025/util"
)

func main() {
	lines := util.GetLinesFromStdin()
	// part1(lines)
	part2(lines)
}

func part2(lines []string) {
	log.Println("part2")

	// debug := false
	debug := true

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

		go part2_worker(s, e, chWorkers, debug)
	}

	// Wait for all workers to finish before continuing
	totalInvalidIDs := 0
	for range len(ranges) {
		invalidIDs := <-chWorkers
		totalInvalidIDs += invalidIDs
	}

	log.Printf("%d", totalInvalidIDs)
}

func part2_worker(s int, e int, chWorkers chan int, debug bool) {
	invalidIDs := 0

	// Iterate over range of IDs to find a ids which consists ONLY of a sequence twice
	for i := s; i <= e; i++ {
		curr := strconv.Itoa(i)

		// Find sequence which repeats atleast twice
		seqTemplate := ""
		seqIndex := 0
		sequences := []string{""}
		latestSeqIndex := 0

		fnAddToTemplate := func(b byte) {
			seqTemplate = seqTemplate + string(b)
		}

		fnAddToLatestSeq := func(b byte) {
			sequences[latestSeqIndex] = sequences[latestSeqIndex] + string(b)
		}

		// pseudocode:
		// 121412141214
		// 121
		// FOR each rune
		//  IS this rune a match for the current template?
		//   ADD to latest seq
		//  ELSE if there is nothing left in this template to check?
		//   ADD rune to latest seq
		//   ADD rune to template
		//  ELSE
		//   Squash all sequences back into a single seq
		//   Set seq tempalte and index accordingly
		//   ADD rune to latest seq
		//   ADD rune to template
		for j := range len(curr) {
			latestSeqIndex = len(sequences) - 1

			b := curr[j]

			// IF b matches the start of the template, start a new sequence
			if len(seqTemplate) > seqIndex && seqTemplate[seqIndex] == b {
				// IF b matches the template, ... continue building our latest sequence
				fnAddToLatestSeq(b)
				seqIndex++
			} else if len(sequences) > 1 {
				// Otherwise we must squash all our sequences into one template continue from there
				seqTemplate = strings.Join(sequences, "")
				latestSeqIndex = 0
				seqIndex = 0
				fnAddToTemplate(b)
				fnAddToLatestSeq(b)
			} else if len(seqTemplate) > 0 && seqTemplate[0] == b {
				sequences = append(sequences, "")
				latestSeqIndex++
				fnAddToLatestSeq(b)
				seqIndex = 1
			} else {
				// If the template does not yet have something, for this index, we add to the tempalte
				fnAddToTemplate(b)
				fnAddToLatestSeq(b)
			}
		}

		if len(sequences) >= 2 {
			// Now we know it must be an invalid ID
			if debug {
				log.Printf("curr: '%s'\n", curr)
				log.Printf("%+v", sequences)
			}
			invalidIDs += i
		}
	}

	chWorkers <- invalidIDs
}

func part1(lines []string) {
	debug := false
	// debug := true

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

		go part1_worker(s, e, chWorkers, debug)
	}

	// Wait for all workers to finish before continuing
	totalInvalidIDs := 0
	for range len(ranges) {
		invalidIDs := <-chWorkers
		totalInvalidIDs += invalidIDs
	}

	log.Printf("%d", totalInvalidIDs)
}

func part1_worker(s int, e int, chWorkers chan int, debug bool) {
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
			if debug {
				log.Printf("curr: '%s'\n", curr)
			}
			invalidIDs += i
		}
	}

	chWorkers <- invalidIDs
}
