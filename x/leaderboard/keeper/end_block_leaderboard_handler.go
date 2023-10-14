package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/perfogic/cosmos-checkers/x/leaderboard/types"
)

func (k Keeper) CollectSortAndClipLeaderboard(ctx sdk.Context) {
	leaderboard := k.GetLeaderboard(ctx)
	fmt.Println("Touch here")
	updated := types.AddCandidatesAtNow(leaderboard.Winners, ctx.BlockTime(), k.GetAllCandidates(ctx))
	params := k.GetParams(ctx)
	if params.Length < uint64(len(updated)) {
		updated = updated[:params.Length]
	}
	leaderboard.Winners = updated
	k.SetLeaderboard(ctx, leaderboard)
}
