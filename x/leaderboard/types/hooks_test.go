package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/perfogic/cosmos-checkers/x/cosmoscheckers/testutil"
	"github.com/perfogic/cosmos-checkers/x/cosmoscheckers/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestMultiHookCallsThem(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	hook1 := testutil.NewMockCosmoscheckersHooks(ctrl)
	hook2 := testutil.NewMockCosmoscheckersHooks(ctrl)
	call1 := hook1.EXPECT().AfterPlayerInfoChanged(gomock.Any(), types.PlayerInfo{
		Index:          "alice",
		WonCount:       1,
		LostCount:      2,
		ForfeitedCount: 3,
	}).Times(1)
	hook2.EXPECT().AfterPlayerInfoChanged(gomock.Any(), types.PlayerInfo{
		Index:          "alice",
		WonCount:       1,
		LostCount:      2,
		ForfeitedCount: 3,
	}).Times(1).After(call1)

	multi := types.NewMultiCosmoscheckersHooks(hook1, hook2)
	multi.AfterPlayerInfoChanged(sdk.NewContext(nil, tmproto.Header{}, false, nil), types.PlayerInfo{
		Index:          "alice",
		WonCount:       1,
		LostCount:      2,
		ForfeitedCount: 3,
	})
}
