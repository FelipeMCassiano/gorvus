package subcommands

import (
	"fmt"
	"os"

	kubebuilders "github.com/FelipeMCassiano/gorvus/internal/builders/kube_builders"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

func CreateKubeNewCommand() *cobra.Command {
	kubenewcmd := &cobra.Command{
		Use:   "new",
		Short: "Creates a new kubernetes deployment yaml file",
		Run: func(cmd *cobra.Command, args []string) {
			if err := kubebuilders.DeploymentKubeBuilder(); err != nil {
				fmt.Println(text.FgRed.Sprint(err))
				os.Exit(1)
			}

			fmt.Println(text.FgGreen.Sprint("\nkubenertes deployment created succesfully!"))
		},
	}

	return kubenewcmd
}
