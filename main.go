package main

import (
	"os"

	"github.com/florinutz/go-donut-challenge/pkg"

	"github.com/florinutz/go-donut-challenge/cmds"

	"github.com/florinutz/go-donut-challenge/pkg/app"
	"github.com/spf13/cobra"
)

func main() {
	cmdName := os.Args[0]
	if err := buildRootCommand(cmdName).Execute(); err != nil {
		os.Exit(1)
	}
}

func buildRootCommand(name string) *cobra.Command {
	a := app.New(name)

	cmd := &cobra.Command{
		Use:     a.Name(),
		Short:   "a small cli app client that talks to 2 coinbase api endpoints",
		Version: pkg.Version,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			a.Out = cmd.OutOrStdout()
			a.Err = cmd.OutOrStderr()
		},
	}

	a.ExtendVersionTemplate(cmd, pkg.Commit, pkg.BuildTime)

	cmd.PersistentFlags().BoolVar(&a.Debug, "debug", false, "enables debug mode")

	cmd.AddCommand(
		cmds.BuildCompletionCmd(a),
		cmds.BuildTickerCmd(a),
		cmds.BuildProductsCmd(a),
		cmds.BuildOrderCmd(a),
	)

	return cmd
}
