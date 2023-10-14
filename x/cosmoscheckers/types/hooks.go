package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ CosmoscheckersHooks = MultiCosmoscheckersHooks{}

type MultiCosmoscheckersHooks []CosmoscheckersHooks

func NewMultiCosmoscheckersHooks(hooks ...CosmoscheckersHooks) MultiCosmoscheckersHooks {
	return hooks
}

func (h MultiCosmoscheckersHooks) AfterPlayerInfoChanged(ctx sdk.Context, playerInfo PlayerInfo) {
	for i := range h {
		h[i].AfterPlayerInfoChanged(ctx, playerInfo)
	}
}
