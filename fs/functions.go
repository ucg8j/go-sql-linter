package fs

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// https://cloud.google.com/bigquery/docs/reference/standard-sql/lexical#reserved_keywords
var reservedKeywords = []string{
	"ALL", "AND", "ANY", "ARRAY", "AS", "ASC", "ASSERT_ROWS_MODIFIED", "AT", "BETWEEN", "BY", "CASE", "CAST", "COLLATE", "CONTAINS", "CREATE", "CROSS", "CUBE", "CURRENT", "DEFAULT", "DEFINE", "DESC", "DISTINCT", "ELSE", "END", "ENUM", "ESCAPE", "EXCEPT", "EXCLUDE", "EXISTS", "EXTRACT", "FALSE", "FETCH", "FOLLOWING", "FOR", "FROM", "FULL", "GROUP", "GROUPING", "GROUPS", "HASH", "HAVING", "IF", "IGNORE", "IN", "INNER", "INTERSECT", "INTERVAL", "INTO", "IS", "JOIN", "LATERAL", "LEFT", "LIKE", "LIMIT", "LOOKUP", "MERGE", "NATURAL", "NEW", "NO", "NOT", "NULL", "NULLS", "OF", "ON", "OR", "ORDER", "OUTER", "OVER", "PARTITION", "PRECEDING", "PROTO", "RANGE", "RECURSIVE", "RESPECT", "RIGHT", "ROLLUP", "ROWS", "SELECT", "SET", "SOME", "STRUCT", "TABLESAMPLE", "THEN", "TO", "TREAT", "TRUE", "UNBOUNDED", "UNION", "UNNEST", "USING", "WHEN", "WHERE", "WINDOW", "WITH", "WITHIN"}

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

func CapitaliseKeywords(lines []string, lint bool) []string {
	// init store of lines that break rules
	offendingLines := []string{}
	newLines := []string{}

	for index, line := range lines {
		// Rudimentary approach without lexical parsing, trim the start of the line and skip if it is a comment or jinja syntax
		lineTrim := strings.TrimLeft(line, " ")
		if strings.HasPrefix(lineTrim, "--") || strings.HasPrefix(lineTrim, "#") || strings.HasPrefix(lineTrim, "{{") || strings.HasPrefix(lineTrim, "\"") || strings.HasPrefix(lineTrim, "/*") {
			newLines = append(newLines, line)
			continue
		}

		newLine := []string{}
		comment := false
		// TODO make the following two loops O(n), currently O(n^2)
		for _, word := range strings.Split(line, " ") {

			// check for comment
			if word == "--" || word == "#" {
				comment = true
				newLine = append(newLine, word)
				continue
			} else if comment {
				newLine = append(newLine, word)
				continue
			}

			// check for keywords
			isKeyword := false
			for _, keyword := range reservedKeywords {
				if word == keyword && !comment {
					isKeyword = true
					newLine = append(newLine, word)
					break
				} else if strings.ToUpper(word) == keyword && !comment {
					isKeyword = true
					offendingLines = append(offendingLines, fmt.Sprintf("line %v, issue = %v Keyword not capitalised", index+1, word))
					newLine = append(newLine, strings.ToUpper(word))
					break
				}
			}

			// else append line
			if !isKeyword && !comment {
				newLine = append(newLine, word)
			}
		}
		newLines = append(newLines, strings.Join(newLine, " "))
	}

	if lint {
		return offendingLines
	}
	return newLines
}
