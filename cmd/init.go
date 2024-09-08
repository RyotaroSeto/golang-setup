/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"goup/internal/executil"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [module name]",
	Short: "Initialize a new Go module",
	Long: `Initialize a new Go module with the specified module name.
This will run 'go mod init', 'go mod tidy', and create a main.go file.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		moduleName := args[0]
		fmt.Printf("Initializing Go module: %s\n", moduleName)

		if err := executil.RunCommand("go", "mod", "init", moduleName); err != nil {
			return fmt.Errorf("failed to run go mod init: %w", err)
		}

		if err := executil.RunCommand("go", "mod", "tidy"); err != nil {
			return fmt.Errorf("failed to run go mod tidy: %w", err)
		}

		if err := createMainGo(); err != nil {
			return fmt.Errorf("failed to create main.go: %w", err)
		}

		fmt.Println("Go module initialization complete.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func createMainGo() error {
	mainFileContent := `package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
`
	err := os.WriteFile("main.go", []byte(mainFileContent), 0644)
	if err != nil {
		return err
	}
	fmt.Println("Created main.go")
	return nil
}
