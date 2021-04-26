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

		// print lint failures to console
		for _, line := range offendingLines {
			fmt.Println(line)
		}
	},
}
