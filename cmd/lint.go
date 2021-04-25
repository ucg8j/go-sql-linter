package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(lintCommand)
}

var lintCommand = &cobra.Command{
	Use:   "lint [model name]",
	Short: "Lints the SQL",
	Args:  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// read lines
		lines := LinesInFile(args[0])

		// init store of lines that break rules
		offendingLines := []string{}

		// run rule checker functions
		offendingLines = trailingWhitespace(lines, offendingLines)

		// print lint failures to console
		for _, line := range offendingLines {
			fmt.Println(line)
		}
	},
}

func LinesInFile(filename string) []string {
	f, _ := os.Open(filename)
	scanner := bufio.NewScanner(f)
	result := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result
}

func trailingWhitespace(lines []string, offendingLines []string) []string {
	for index, line := range lines {
		newline := strings.TrimRight(line, " ")

		if line != newline {
			offendingLines = append(offendingLines, fmt.Sprintf("line %v, issue = Trailing whitespace", index))
		}
	}
	return offendingLines
}
