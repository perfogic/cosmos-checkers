package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	checkerstypes "github.com/perfogic/cosmos-checkers/x/cosmoscheckers/types"
	"github.com/perfogic/cosmos-checkers/x/leaderboard/types"
)

var _ checkerstypes.CosmoscheckersHooks = Hooks{}

func (h Hooks) AfterPlayerInfoChanged(ctx sdk.Context, playerInfo checkerstypes.PlayerInfo) {
	candidate, err := types.MakeCandidateFromPlayerInfo(playerInfo)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	if candidate.WonCount < 1 {
		return
	}
	h.k.SetCandidate(ctx, candidate)
}
