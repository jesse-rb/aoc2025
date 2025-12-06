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
