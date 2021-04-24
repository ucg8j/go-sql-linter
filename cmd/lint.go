package cmd

import (
	"fmt"
	"os"
	"bufio"
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
		// trimWhiteSpace(file)
		lines := LinesInFile(args[0])

		offendingLines := map[int][]string{}

		for index, line := range lines  {
			newline := strings.TrimRight(line, " ")

			if line != newline {
				offendingLines[index] = append(offendingLines[index], "Trailing whitespace")
			}
		}
		for index, line := range offendingLines {
			fmt.Printf("line %v, issue = %v\n", index , line)
		}
	},
}

func LinesInFile(filename string) []string  {
	f, _ := os.Open(filename)
	scanner := bufio.NewScanner(f)
	result := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result
}
