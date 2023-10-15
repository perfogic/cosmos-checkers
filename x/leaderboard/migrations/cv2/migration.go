package cv2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	cv2keeper "github.com/perfogic/cosmos-checkers/x/leaderboard/migrations/cv2/keeper"
	cv2types "github.com/perfogic/cosmos-checkers/x/leaderboard/migrations/cv2/types"
	"github.com/perfogic/cosmos-checkers/x/leaderboard/types"
)

func ComputeMigratedLeaderboard(ctx sdk.Context, playerInfosk types.PlayerInfoKeeper) (*types.Leaderboard, error) {
	return cv2keeper.MapPlayerInfosReduceToLeaderboard(
		sdk.WrapSDKContext(ctx),
		playerInfosk,
		types.DefaultLength,
		ctx.BlockTime(),
		cv2types.PlayerInfoChunkSize)
}

func ComputeInitGenesis(ctx sdk.Context, playerInfosk types.PlayerInfoKeeper) (*types.GenesisState, error) {
	leaderboard, err := ComputeMigratedLeaderboard(ctx, playerInfosk)
	if err != nil {
		return nil, err
	}
	return &types.GenesisState{
		Params:      types.DefaultParams(),
		Leaderboard: *leaderboard,
	}, nil
}
