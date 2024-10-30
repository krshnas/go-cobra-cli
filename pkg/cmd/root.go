package cmd

import (
	"fmt"
	"os"

	filem "github.com/krishna/cobra-cli/pkg/cmd/filemanipulation"
	substr "github.com/krishna/cobra-cli/pkg/cmd/stringmanipulation"
	"github.com/spf13/cobra"
)

const version = "1.0"

var rootCmd = &cobra.Command{
	Use:     "stringutil",
	Version: version,
	Short:   "A Simple String Utility",
	Long:    `String Utility is a Cobra framework-based CLI application that provides simple text manipulation functions.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is a Cobra based String Utility.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(substr.NewCmdSubstr())
	rootCmd.AddCommand(filem.FileCmd)
}
