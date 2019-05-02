package main

import (
	"github.com/radu-matei/coras/pkg/coras"

	"github.com/spf13/cobra"
)

type pushCmd struct {
	inputFile string
	targetRef string
	exported  bool
}

func newPushCmd() *cobra.Command {
	const usage = `pushes a CNAB bundle to a registry using ORAS`
	var p pushCmd

	cmd := &cobra.Command{
		Use:   "push",
		Short: usage,
		Long:  usage,
		RunE: func(cmd *cobra.Command, args []string) error {
			p.inputFile = args[0]
			p.targetRef = args[1]
			return p.run()
		},
	}
	cmd.Flags().BoolVarP(&p.exported, "exported", "", false, "When passed, this command will push an exported (thick) bundle")

	return cmd
}

func (p *pushCmd) run() error {
	return coras.Push(p.inputFile, p.targetRef, p.exported)
}
