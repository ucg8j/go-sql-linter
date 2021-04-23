package cmd

import (
	"fmt"

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
		fmt.Println("hello lint")
	},
	// ValidArgsFunction: completeModelFn,
}
