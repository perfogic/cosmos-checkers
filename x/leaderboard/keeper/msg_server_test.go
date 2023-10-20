package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/perfogic/cosmos-checkers/testutil/keeper"
	"github.com/perfogic/cosmos-checkers/x/leaderboard/keeper"
	"github.com/perfogic/cosmos-checkers/x/leaderboard/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.LeaderboardKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
