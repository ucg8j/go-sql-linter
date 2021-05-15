package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"go-sql-linter/fs"
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
		lines := fs.ReadLinesInFile(args[0])
		lint := true

		// run rule checker functions
		offendingLines := fs.TrailingWhitespace(lines, lint)
		// TODO make elegant
		offendingLinesTmp := fs.MultipleNewLines(lines, lint)
		offendingLines = append(offendingLines, offendingLinesTmp...)
		offendingLinesTmp = fs.CapitaliseKeywords(lines, lint)
		offendingLines = append(offendingLines, offendingLinesTmp...)

		// print lint failures to console
		for _, line := range offendingLines {
			fmt.Println(line)
		}
	},
}
