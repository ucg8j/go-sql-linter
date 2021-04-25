package cmd

import (
	"fmt"
	"os"
	"bufio"

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
		lines := fs.LinesInFile(args[0])
		lint := false

		// fix lines
		lines = fs.TrailingWhitespace(lines, lint)

		// write new file
		fmt.Println("Writing new file...")
		file, err := os.Create("./temp.sql")
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
	},
}
