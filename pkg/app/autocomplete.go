package app

import (
	"io"

	"github.com/florinutz/go-donut-challenge/pkg/config"

	"github.com/spf13/cobra"
)

func (app *App) AutoComplete(rootCmd *cobra.Command, shellName string) error {
	var shell config.Shell

	shell, err := config.ParseShell(shellName)
	if err != nil {
		return err
	}

	completionFuncs := map[config.Shell]func(io.Writer) error{
		config.ShellBash: rootCmd.GenBashCompletion,
		config.ShellZsh:  rootCmd.GenZshCompletion,
	}

	return completionFuncs[shell](app.Out)
}
