package cmd

import (
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "fiber-crud-generator",
    Short: "A CLI tool to generate CRUD resources for Go Fiber with MongoDB",
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        panic(err)
    }
}