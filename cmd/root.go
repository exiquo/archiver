package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "archiver",
	Short: "archiver is a cli tool for archiving directories",
	Long:  "archiver is a cli tool for archiving directories - it's creating a .tar.gz file at a target directory",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if err := ArchiveDirectory(args[0], args[1]); err != nil {
			fmt.Fprintf(os.Stderr, "an error occurred while executing archiver '%s'\n", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "an error occurred while executing archiver '%s'\n", err)
		os.Exit(1)
	}
}
