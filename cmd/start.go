package cmd

import (
	"github.com/bnema/floyd/internal/app"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start [project]",
	Short: "Start a Flowmodoro session for a project",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var project string
		if len(args) > 0 {
			project = args[0]
		}
		return app.Run(project)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
