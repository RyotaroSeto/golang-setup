/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"goup/internal/executil"

	"github.com/spf13/cobra"
)

var benchmarkCmd = &cobra.Command{
	Use:   "benchmark",
	Short: "Run benchmark tests in the project",
	Long: `The goup benchmark command runs all benchmark tests in the current Go project.
You can specify specific benchmarks or run all benchmarks in the package.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		benchPattern, _ := cmd.Flags().GetString("bench")

		testArgs := []string{"test", "-bench=.", "-run=^$", "./..."}

		if benchPattern != "" {
			testArgs = []string{"test", fmt.Sprintf("-bench=%s", benchPattern), "-run=^$", "./..."}
		}

		if err := executil.RunGoCommand(testArgs...); err != nil {
			return fmt.Errorf("failed to run go benchmark: %w", err)
		}

		return nil
	},
}

func init() {
	benchmarkCmd.Flags().StringP("bench", "b", "", "Specify the benchmark function to run (e.g. 'BenchmarkXxx')")
	rootCmd.AddCommand(benchmarkCmd)
}
