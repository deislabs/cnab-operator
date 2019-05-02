package main

import (
	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	const usage = `Push and pull CNAB bundles using ORAS`

	cmd := &cobra.Command{
		Use:          "coras",
		Short:        usage,
		SilenceUsage: true,
	}

	cmd.AddCommand(
		newPushCmd(),
		newPullCmd(),
	)

	return cmd
}
