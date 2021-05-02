package fs

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadLinesInFile(filename string) []string {
	if filepath.Ext(strings.TrimSpace(filename)) != ".sql" {
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

func WriteLinesInFile(filename string, lines []string) {
	// write new file
	fmt.Println("Writing new file...")
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			panic("Got error while writing to a file")
		}
	}
	writer.Flush()
}

func TrailingWhitespace(lines []string, lint bool) []string {
	// init store of lines that break rules
	offendingLines := []string{}
	newLines := []string{}

	for index, line := range lines {
		newline := strings.TrimRight(line, " ")

		if line != newline && lint {
			offendingLines = append(offendingLines, fmt.Sprintf("line %v, issue = Trailing whitespace", index))
		} else {
			newLines = append(newLines, newline)
		}
	}

	if lint {
		return offendingLines
	}
	return newLines

}

func MultipleNewLines(lines []string, lint bool) []string {
	// init store of lines that break rules
	offendingLines := []string{}
	newLines := []string{}

	var previousLine string

	for index, line := range lines {
		// not possible to have consecutive newlines
		if index == 0 {
			previousLine = line
			newLines = append(newLines, line)
			continue
		}

		// newlines aren't represented by '\n' because the scanner in ReadLinesInFile uses it as the separator
		if previousLine == line && line == "" && lint {
			offendingLines = append(offendingLines, fmt.Sprintf("line %v, issue = Multiple new lines", index))
		} else if previousLine == line && line == "" && (len(lines)-1) != index {
			// no need to append if empty line
			continue
		} else {
			newLines = append(newLines, line)
		}

		previousLine = line
	}

	if lint {
		return offendingLines
	}
	return newLines

}
