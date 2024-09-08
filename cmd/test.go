/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"goup/internal/executil"

	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run all tests in the project",
	Long:  `The goup test command runs all tests in the current Go project. It supports optional flags for verbose output and coverage reports.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		testArgs := []string{"test", "./...", "-v", "-cover"}

		if err := executil.RunGoCommand(testArgs...); err != nil {
			fmt.Errorf("failed to run go test: %w", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
