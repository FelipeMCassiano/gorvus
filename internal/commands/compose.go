package commands

import (
	"github.com/FelipeMCassiano/gorvus/internal/commands/subcommands"
	"github.com/spf13/cobra"
)

func CreateComposeCommand() *cobra.Command {
	composeCmd := &cobra.Command{
		Use:   "compose",
		Short: "Manages current directory's docker-compose.yml",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	composeCmd.AddCommand(subcommands.CreateComposeAddCommand())
	composeCmd.AddCommand(subcommands.CreateComposeAddNetCommand())
	composeCmd.AddCommand(subcommands.CreateComposeCreateCommand())
	composeCmd.AddCommand(subcommands.CreateComposeRemoveCommand())

	return composeCmd
}
