package main

import (
	_ "embed"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{Use: "geosend"}
	rootCmd.AddCommand(SendCmd)
	rootCmd.AddCommand(ReceiveCmd)
	rootCmd.AddCommand(RelayCmd) // <-- Add this line
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
