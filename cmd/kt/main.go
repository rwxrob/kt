package main

import (
	"os"

	kt "github.com/rwxrob/kt/pkg/cmd"
	"github.com/spf13/cobra"
)

// completion has to be here because needs to see kt.Cmd
// and would create a circular dependency if in pkg
var completionCmd = &cobra.Command{
	Use:       "completion (bash|zsh|fish|pwsh)",
	Short:     "Generate completion script",
	Args:      cobra.MatchAll(cobra.ExactValidArgs(1)),
	ValidArgs: []string{`bash`, `zsh`, `fish`, `pwsh`},
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			kt.Cmd.GenBashCompletion(os.Stdout)
		case "zsh":
			kt.Cmd.GenZshCompletion(os.Stdout)
		case "fish":
			kt.Cmd.GenFishCompletion(os.Stdout, true)
		case "pwsh":
			kt.Cmd.GenPowerShellCompletion(os.Stdout)
		}
	},
}

func main() {
	kt.Cmd.AddCommand(completionCmd)
	kt.Cmd.Execute()
}
