package cmd

import (
	"github.com/bnema/floyd/internal/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "floyd",
	Short:   "Floyd is a Flowmodoro timer application",
	Long:    `Floyd helps you manage your work and break cycles using the Flowmodoro technique.`,
	Version: version.PrintVersion(),
}

func Execute() error {
	return rootCmd.Execute()
}
