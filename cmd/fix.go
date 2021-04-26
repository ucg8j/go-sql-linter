package cmd

import (
	"github.com/spf13/cobra"

	"go-sql-linter/fs"
)

func init() {
	rootCmd.AddCommand(fixCommand)
}

var fixCommand = &cobra.Command{
	Use:   "fix [model name]",
	Short: "Fixes lint issues in a SQL",
	Args:  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// read lines
		lines := fs.ReadLinesInFile(args[0])
		lint := false

		// fix lines
		lines = fs.TrailingWhitespace(lines, lint)

		// write new file
		fs.WriteLinesInFile(args[0], lines)
	},
}
