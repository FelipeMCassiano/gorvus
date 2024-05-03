package commands

import (
	"fmt"
	"os"

	"github.com/FelipeMCassiano/gorvus/internal/commands/subcommands"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

func CreateKubeCommand() *cobra.Command {
	kubecmd := &cobra.Command{
		Use:   "kube",
		Short: "Manages yaml kubernetes files",
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				fmt.Println(text.FgRed.Sprint(err))
				os.Exit(1)
			}
		},
	}
	kubecmd.AddCommand(subcommands.CreateKubeDeploymentCommand())
	return kubecmd
}
