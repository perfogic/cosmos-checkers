package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/perfogic/cosmos-checkers/testutil/keeper"
	"github.com/perfogic/cosmos-checkers/x/cosmoscheckers/keeper"
	"github.com/perfogic/cosmos-checkers/x/cosmoscheckers/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CosmoscheckersKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestCreateGame(t *testing.T) {
	msgServer, context := setupMsgServer(t)
	createResponse, err := msgServer.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgCreateGameResponse{
		GameIndex: "First Test It", // TODO: update with a proper value when updated
	}, *createResponse)
}
