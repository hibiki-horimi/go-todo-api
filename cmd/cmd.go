package cmd

import "github.com/spf13/cobra"

func Register(rootCmd *cobra.Command) {
	rootCmd.AddCommand(
		serverCmd,
		migrateCmd,
	)
}
