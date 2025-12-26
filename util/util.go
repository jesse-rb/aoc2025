// Package util is a small utiity package which all days of aoc2025 can use
package util

import (
	"bufio"
	"log"
	"os"
)

func GetLinesFromStdin() []string {
	lines := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func Must(err error) {
	if err != nil {
		log.Fatalf("An unexpected error occured")
	}
}

func AbsInt(x int) int {
	if x < 0 {
		return (-1) * x
	} else {
		return x
	}
}

func InverseSignInt(x int) int {
	return -1 * x
}

func ModInt(a int, b int) int {
	return (a%b + b) % b
}

func IfThenElse[T any](expression bool, v1 T, v2 T) T {
	if expression {
		return v1
	} else {
		return v2
	}
}

func CeilInt(a int, b int) int {
	return (a + b - 1) / b
}

func MaxInt64(a int64, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}
