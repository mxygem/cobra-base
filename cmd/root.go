package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	debug = false
	dry   = false

	rootCmd = &cobra.Command{
		Use:   "base",
		Short: "add a short description here",
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "turn on debug mode")
	rootCmd.PersistentFlags().BoolVar(&dry, "dry", false, "go through steps, but do not apply changes")
}

func Execute() error {
	return rootCmd.Execute()
}

func requireArgs(count int) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		if len(args) >= count {
			return
		}

		fmt.Printf("command %s requires %d arguments, received %d.\n", cmd.Name(), count, len(args))
		if err := cmd.Usage(); err != nil {
			fmt.Printf("problem calling Usage for command: %v\n", err)
		}
		os.Exit(1)
	}
}
