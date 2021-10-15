package cmd

import (
	"strconv"

	"github.com/mxygem/cobra-base/internal/add"

	"github.com/spf13/cobra"
)

var (
	addCmd = &cobra.Command{
		Use:    "add [x] [y]",
		Short:  "return the sum of two provided numbers",
		PreRun: requireArgs(2),
		RunE:   addRun,
	}
)

func init() {
	rootCmd.AddCommand(addCmd)
}

func addRun(cmd *cobra.Command, args []string) error {
	// handle arg prep
	// in other situations this could be env vars, setting up logging, etc.
	// *IMPORTANT* - the goal here should be to do the MINIMAL amount of work here in order to call
	// the internal function whose purpose is to actually do the work needed
	x, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return err
	}

	y, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return err
	}

	add.Run(x, y)

	return nil
}
