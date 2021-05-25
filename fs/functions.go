package fs

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// https://cloud.google.com/bigquery/docs/reference/standard-sql/lexical#reserved_keywords
var reservedKeywords = map[string]int{
	"ALL":                  0,
	"AND":                  1,
	"ANY":                  2,
	"ARRAY":                3,
	"AS":                   4,
	"ASC":                  5,
	"ASSERT_ROWS_MODIFIED": 6,
	"AT":                   7,
	"BETWEEN":              8,
	"BY":                   9,
	"CASE":                 10,
	"CAST":                 11,
	"COLLATE":              12,
	"CONTAINS":             13,
	"CREATE":               14,
	"CROSS":                15,
	"CUBE":                 16,
	"CURRENT":              17,
	"DEFAULT":              18,
	"DEFINE":               19,
	"DESC":                 20,
	"DISTINCT":             21,
	"ELSE":                 22,
	"END":                  23,
	"ENUM":                 24,
	"ESCAPE":               25,
	"EXCEPT":               26,
	"EXCLUDE":              27,
	"EXISTS":               28,
	"EXTRACT":              29,
	"FALSE":                30,
	"FETCH":                31,
	"FOLLOWING":            32,
	"FOR":                  33,
	"FROM":                 34,
	"FULL":                 35,
	"GROUP":                36,
	"GROUPING":             37,
	"GROUPS":               38,
	"HASH":                 39,
	"HAVING":               40,
	"IF":                   41,
	"IGNORE":               42,
	"IN":                   43,
	"INNER":                44,
	"INTERSECT":            45,
	"INTERVAL":             46,
	"INTO":                 47,
	"IS":                   48,
	"JOIN":                 49,
	"LATERAL":              50,
	"LEFT":                 51,
	"LIKE":                 52,
	"LIMIT":                53,
	"LOOKUP":               54,
	"MERGE":                55,
	"NATURAL":              56,
	"NEW":                  57,
	"NO":                   58,
	"NOT":                  59,
	"NULL":                 60,
	"NULLS":                61,
	"OF":                   62,
	"ON":                   63,
	"OR":                   64,
	"ORDER":                65,
	"OUTER":                66,
	"OVER":                 67,
	"PARTITION":            68,
	"PRECEDING":            69,
	"PROTO":                70,
	"RANGE":                71,
	"RECURSIVE":            72,
	"RESPECT":              73,
	"RIGHT":                74,
	"ROLLUP":               75,
	"ROWS":                 76,
	"SELECT":               77,
	"SET":                  78,
	"SOME":                 79,
	"STRUCT":               80,
	"TABLESAMPLE":          81,
	"THEN":                 82,
	"TO":                   83,
	"TREAT":                84,
	"TRUE":                 85,
	"UNBOUNDED":            86,
	"UNION":                87,
	"UNNEST":               88,
	"USING":                89,
	"WHEN":                 90,
	"WHERE":                91,
	"WINDOW":               92,
	"WITH":                 93,
	"WITHIN":               94,
}

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
			offendingLines = append(offendingLines, fmt.Sprintf("line %v, issue = Trailing whitespace", index+1))
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

		for _, word := range strings.Split(line, " ") {

			wordUpper := strings.ToUpper(word)
			wordTrim := strings.Trim(wordUpper, ",()")

			// check for inline comments
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

			if _, found := reservedKeywords[wordTrim]; found && !comment {
				isKeyword = true
				offendingLines = append(offendingLines, fmt.Sprintf("line %v, issue = %v Keyword not capitalised", index+1, wordTrim))
				newLine = append(newLine, wordUpper)
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
