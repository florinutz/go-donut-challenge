package cmds

import (
	"github.com/florinutz/go-donut-challenge/pkg/app"
	"github.com/florinutz/go-donut-challenge/pkg/config"
	"github.com/spf13/cobra"
)

func BuildCompletionCmd(app *app.App) *cobra.Command {
	return &cobra.Command{
		Use:     "completion <bash|zsh>",
		Aliases: []string{"autocomplete"},
		Short:   "Generates shell completion scripts",
		Long: `
To load completion for bash run
	. <(go-donut-challenge completion bash)

To load for zsh run 
	. <(go-donut-challenge completion zsh)

To configure your shell to load completions for each session add this to your bash's .bashrc or zsh's .zshrc:
	. <(go-donut-challenge completion bash)

To configure your zsh shell to load completions for each session add this to your .zshrc:
	. <(go-donut-challenge completion zsh)

And then source them:
	source ~/.zshrc

or restart the terminal.
`,
		Args:      cobra.ExactArgs(1),
		ValidArgs: []string{string(config.ShellBash), string(config.ShellZsh)},
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.AutoComplete(cmd.Root(), args[0])
		},
	}
}
