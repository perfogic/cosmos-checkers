package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/perfogic/cosmos-checkers/x/cosmoscheckers/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateGame() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-game [black] [red] [wager] [denom]",
		Short: "Broadcast message createGame",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argBlack := args[0]
			argRed := args[1]
			argWager, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}
			argDenom := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateGame(
				clientCtx.GetFromAddress().String(),
				argBlack,
				argRed,
				argWager,
				argDenom,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}