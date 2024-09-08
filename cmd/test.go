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
		testName, _ := cmd.Flags().GetString("run")

		testArgs := []string{"test", "./...", "-v", "-cover"}

		if testName != "" {
			testArgs = append(testArgs, fmt.Sprintf("-run=%s", testName))
		}

		if err := executil.RunGoCommand(testArgs...); err != nil {
			fmt.Errorf("failed to run go test: %w", err)
		}
		return nil
	},
}

func init() {
	testCmd.Flags().StringP("run", "r", "", "Specify a test function to run (e.g., 'TestXxx')")
	rootCmd.AddCommand(testCmd)
}
