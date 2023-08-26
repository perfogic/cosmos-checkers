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

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CosmoscheckersKeeper(t)
	cosmoscheckers.InitGenesis(ctx, *k, genesisState)
	got := cosmoscheckers.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
