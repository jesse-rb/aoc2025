package util

import (
	"bufio"
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
