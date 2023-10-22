package cv3

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/perfogic/cosmos-checkers/x/cosmoscheckers/keeper"
)

func PerformMigration(ctx sdk.Context, k keeper.Keeper, storedGameChunk uint64) error {
	ctx.Logger().Info("Start to compute checkers games to player info calculation...")
	return nil
}
