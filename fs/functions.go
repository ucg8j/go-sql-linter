package fs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LinesInFile(filename string) []string {
	if strings.TrimSpace(filename) != ".sql" {
		fmt.Println("❌ Please provide a file with the .sql extension")
		os.Exit(1)
	}
	f, _ := os.Open(filename)
	scanner := bufio.NewScanner(f)
	result := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result
}

func TrailingWhitespace(lines []string, lint bool) []string {
	// init store of lines that break rules
	offendingLines := []string{}
	newLines := []string{}

	for index, line := range lines {
		newline := strings.TrimRight(line, " ")

		if line != newline && lint {
			offendingLines = append(offendingLines, fmt.Sprintf("line %v, issue = Trailing whitespace", index))
		} else if line != newline && !lint {
			newLines = append(newLines, newline)
		} else {
			newLines = append(newLines, line)
		}
	}

	if lint {
		return offendingLines
	}
	return newLines

}
