package keeper

import (
	"github.com/perfogic/cosmos-checkers/x/cosmoscheckers/types"
)

var _ types.QueryServer = Keeper{}
