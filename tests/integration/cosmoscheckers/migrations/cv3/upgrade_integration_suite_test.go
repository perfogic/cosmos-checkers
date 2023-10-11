package cv3_test

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	module "github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	checkersapp "github.com/perfogic/cosmos-checkers/app"
	"github.com/perfogic/cosmos-checkers/x/cosmoscheckers/keeper"
	cv2types "github.com/perfogic/cosmos-checkers/x/cosmoscheckers/migrations/cv2/types"
	"github.com/perfogic/cosmos-checkers/x/cosmoscheckers/types"
	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

type IntegrationTestSuite struct {
	suite.Suite

	app         *checkersapp.App
	msgServer   types.MsgServer
	ctx         sdk.Context
	queryClient types.QueryClient
}

func TestUpgradeTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (suite *IntegrationTestSuite) SetupTest() {
	app := checkersapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Time: time.Now()})

	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())
	app.BankKeeper.SetParams(ctx, banktypes.DefaultParams())
	initialVM := module.VersionMap{types.ModuleName: cv2types.ConsensusVersion}
	app.UpgradeKeeper.SetModuleVersionMap(ctx, initialVM)

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.CosmoscheckersKeeper)
	queryClient := types.NewQueryClient(queryHelper)

	suite.app = app
	suite.msgServer = keeper.NewMsgServerImpl(app.CosmoscheckersKeeper)
	suite.ctx = ctx
	suite.queryClient = queryClient
}
