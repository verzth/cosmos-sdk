package cmd

import (
	"github.com/spf13/cobra"
	"github.com/verzth/cosmos-sdk/client"
)

func HomeCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "home",
		Short: "Outputs the string being used as the home path. No home directory is set when using the tool standalone.",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			clientCtx := client.GetClientContextFromCmd(cmd)
			if clientCtx.HomeDir == "" {
				cmd.Println("No home directory set.")
			} else {
				cmd.Println(clientCtx.HomeDir)
			}
		},
	}
}
