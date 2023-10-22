package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/perfogic/cosmos-checkers/x/leaderboard/types"
)

func (k msgServer) UpdateBoard(goCtx context.Context, msg *types.MsgUpdateBoard) (*types.MsgUpdateBoardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	board, found := k.GetBoard(ctx)
	if !found {
		panic(types.ErrBoardNotFound)
	}
	playerInfoList := board.PlayerInfo
	candidates := make([]types.PlayerInfo, 0, len(msg.GetSigners()))
	for _, signer := range msg.GetSigners() {
		candidate, found := k.GetPlayerInfo(ctx, signer.String())
		if found {
			candidates = append(candidates, candidate)
		}
	}
	if len(candidates) == 0 {
		return nil, types.ErrCandidateNotFound
	}
	updated := types.UpdatePlayerInfoList(playerInfoList, candidates)
	board.PlayerInfo = updated
	k.SetBoard(ctx, board)

	return &types.MsgUpdateBoardResponse{}, nil
}
