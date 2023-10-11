package cosmoscheckers_test

import (
	"testing"

	keepertest "github.com/perfogic/cosmos-checkers/testutil/keeper"
	"github.com/perfogic/cosmos-checkers/testutil/nullify"
	"github.com/perfogic/cosmos-checkers/x/cosmoscheckers"
	"github.com/perfogic/cosmos-checkers/x/cosmoscheckers/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		SystemInfo: types.SystemInfo{
			NextId: 36,
		},
		StoredGameList: []types.StoredGame{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		PlayerInfoList: []types.PlayerInfo{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmoscheckersKeeper(t)
	cosmoscheckers.InitGenesis(ctx, *k, genesisState)
	got := cosmoscheckers.ExportGenesis(ctx, *k)
	require.NotNil(t, got)
	require.Equal(t, got.SystemInfo.NextId, uint64(36))
	require.Equal(t, got.StoredGameList[0].Index, "0")
	require.Equal(t, got.StoredGameList[1].Index, "1")

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.SystemInfo, got.SystemInfo)
	require.ElementsMatch(t, genesisState.StoredGameList, got.StoredGameList)
	require.ElementsMatch(t, genesisState.PlayerInfoList, got.PlayerInfoList)
	// this line is used by starport scaffolding # genesis/test/assert
}
