package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gsl",
	Short: "A SQL linter written in go",
	Long:  `gsl is an opinionated SQL linter. Built in Go. gsl aims to be fast and config free`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
