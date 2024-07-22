package commands

import (
	"fmt"
	"os"

	"github.com/FelipeMCassiano/gorvus/internal/commands/subcommands"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

func CreateComposeCommand() *cobra.Command {
	var path string
	composeCmd := &cobra.Command{
		Use:   "compose",
		Short: "Manages current directory's docker-compose.yml",
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				fmt.Println(text.FgRed.Sprint(err))
				os.Exit(1)
			}
		},
	}
	composeCmd.PersistentFlags().StringVar(&path, "cd", "", "Change the Directory")
	composeCmd.AddCommand(subcommands.CreateComposeAddCommand())
	composeCmd.AddCommand(subcommands.CreateComposeAddNetCommand())
	composeCmd.AddCommand(subcommands.CreateComposeNewCommand())
	composeCmd.AddCommand(subcommands.CreateComposeRemoveCommand())

	return composeCmd
}
