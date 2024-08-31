package cmd

import (
	"fmt"

	"github.com/bnema/floyd/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version of Floyd",
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Floyd version %s\n", version.PrintVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
