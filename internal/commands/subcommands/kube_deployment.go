package subcommands

import (
	"fmt"
	"os"

	kubebuilders "github.com/FelipeMCassiano/gorvus/internal/builders/kube_builders"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

func CreateKubeDeploymentCommand() *cobra.Command {
	kubedeploymentcmd := &cobra.Command{
		Use:     "deployment",
		Short:   "Manage kubernetes deployment yaml file",
		Aliases: []string{"depl", "deploy"},
	}
	kubedeploymentcmd.AddCommand(CreateKubeDeploymentNewCommand())

	return kubedeploymentcmd
}

func CreateKubeDeploymentNewCommand() *cobra.Command {
	deplnew := &cobra.Command{
		Use:     "new",
		Short:   "Create a new deployment yaml file",
		Aliases: []string{"n"},
		Run: func(cmd *cobra.Command, args []string) {
			if err := kubebuilders.DeploymentKubeBuilder(); err != nil {
				fmt.Println(text.FgRed.Sprint(err))
				os.Exit(1)
			}

			fmt.Println(text.FgGreen.Sprint("\nkubenertes deployment created succesfully!"))
		},
	}
	return deplnew
}
