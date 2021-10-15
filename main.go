package main

import (
	"fmt"
	"os"

	"github.com/mxygem/cobra-base/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Printf("failed to do a thing: %v", err)
		os.Exit(1)
	}
}
