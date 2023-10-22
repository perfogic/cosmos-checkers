package keeper

import (
	"github.com/perfogic/cosmos-checkers/x/leaderboard/types"
)

var _ types.QueryServer = Keeper{}
