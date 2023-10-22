package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/perfogic/cosmos-checkers/x/leaderboard/types"
)

// SetBoard set board in the store
func (k Keeper) SetBoard(ctx sdk.Context, board types.Board) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BoardKey))
	b := k.cdc.MustMarshal(&board)
	store.Set([]byte{0}, b)
}

// GetBoard returns board
func (k Keeper) GetBoard(ctx sdk.Context) (val types.Board, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BoardKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveBoard removes board from the store
func (k Keeper) RemoveBoard(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BoardKey))
	store.Delete([]byte{0})
}
