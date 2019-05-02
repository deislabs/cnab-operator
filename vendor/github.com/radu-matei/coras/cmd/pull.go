package main

import (
	"github.com/radu-matei/coras/pkg/coras"
	"github.com/spf13/cobra"
)

type pullCmd struct {
	outputBundle string
	targetRef    string
	exported     bool
}

func newPullCmd() *cobra.Command {
	const usage = "pulls a CNAB bundle from a registry using ORAS"
	var p pullCmd

	cmd := &cobra.Command{
		Use:   "pull",
		Short: usage,
		Long:  usage,
		RunE: func(cmd *cobra.Command, args []string) error {
			p.outputBundle = args[0]
			p.targetRef = args[1]
			return p.run()
		},
	}
	cmd.Flags().BoolVarP(&p.exported, "exported", "", false, "When passed, this command will pull an exported (thick) bundle")

	return cmd
}

func (p *pullCmd) run() error {
	return coras.Pull(p.targetRef, p.outputBundle, p.exported)
}
