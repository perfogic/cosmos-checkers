package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/perfogic/cosmos-checkers/x/leaderboard/types"
	"github.com/spf13/cobra"
)

func CmdShowLeaderboard() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-leaderboard",
		Short: "shows leaderboard",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetLeaderboardRequest{}

			res, err := queryClient.Leaderboard(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
