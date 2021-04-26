package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gsl",
	Short: "A SQL linter written in go",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, world")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
