package cv3_test

import (
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/perfogic/cosmos-checkers/app/upgrades/v1tov1_1"
	cv2types "github.com/perfogic/cosmos-checkers/x/cosmoscheckers/migrations/cv2/types"
	cv3types "github.com/perfogic/cosmos-checkers/x/cosmoscheckers/migrations/cv3/types"
	"github.com/perfogic/cosmos-checkers/x/cosmoscheckers/rules"
	"github.com/perfogic/cosmos-checkers/x/cosmoscheckers/testutil"
	"github.com/perfogic/cosmos-checkers/x/cosmoscheckers/types"
)

const (
	alice = testutil.Alice
	bob   = testutil.Bob
	carol = testutil.Carol
)

func (suite *IntegrationTestSuite) TestUpgradeConsensusVersion() {
	vmBefore := suite.app.UpgradeKeeper.GetModuleVersionMap(suite.ctx)
	suite.Require().Equal(cv2types.ConsensusVersion, vmBefore[types.ModuleName])

	v1Tov1_1Plan := upgradetypes.Plan{
		Name:   v1tov1_1.UpgradeName,
		Info:   "some text here",
		Height: 123450000,
	}

	suite.app.UpgradeKeeper.ApplyUpgrade(suite.ctx, v1Tov1_1Plan)

	vmAfter := suite.app.UpgradeKeeper.GetModuleVersionMap(suite.ctx)
	suite.Require().Equal(cv3types.ConsensusVersion, vmAfter[types.ModuleName])
}

func (suite *IntegrationTestSuite) TestNotUpgradeConsensusVersion() {
	vmBefore := suite.app.UpgradeKeeper.GetModuleVersionMap(suite.ctx)
	suite.Require().Equal(cv2types.ConsensusVersion, vmBefore[types.ModuleName])

	dummyPlan := upgradetypes.Plan{
		Name:   v1tov1_1.UpgradeName + "no",
		Info:   "some text here",
		Height: 123450000,
	}
	defer func() {
		r := recover()
		suite.Require().NotNil(r, "The code did not panic")
		suite.Require().Equal(r, "ApplyUpgrade should never be called without first checking HasHandler")

		vmAfter := suite.app.UpgradeKeeper.GetModuleVersionMap(suite.ctx)
		suite.Require().Equal(cv2types.ConsensusVersion, vmAfter[types.ModuleName])
	}()
	suite.app.UpgradeKeeper.ApplyUpgrade(suite.ctx, dummyPlan)
}

func (suite *IntegrationTestSuite) TestUpgradeTallyPlayerInfo() {
	suite.app.CosmoscheckersKeeper.SetStoredGame(suite.ctx, types.StoredGame{
		Index:  "1",
		Black:  alice,
		Red:    bob,
		Winner: rules.PieceStrings[rules.BLACK_PLAYER],
	})
	suite.app.CosmoscheckersKeeper.SetStoredGame(suite.ctx, types.StoredGame{
		Index:  "2",
		Black:  alice,
		Red:    carol,
		Winner: rules.PieceStrings[rules.RED_PLAYER],
	})
	suite.app.CosmoscheckersKeeper.SetStoredGame(suite.ctx, types.StoredGame{
		Index:  "3",
		Black:  alice,
		Red:    carol,
		Winner: rules.PieceStrings[rules.BLACK_PLAYER],
	})
	suite.app.CosmoscheckersKeeper.SetStoredGame(suite.ctx, types.StoredGame{
		Index:  "4",
		Black:  alice,
		Red:    bob,
		Winner: rules.PieceStrings[rules.NO_PLAYER],
	})
	suite.Require().EqualValues([]types.PlayerInfo(nil), suite.app.CosmoscheckersKeeper.GetAllPlayerInfo(suite.ctx))

	v1Tov1_1Plan := upgradetypes.Plan{
		Name:   v1tov1_1.UpgradeName,
		Info:   "some text here",
		Height: 123450000,
	}
	suite.app.UpgradeKeeper.ApplyUpgrade(suite.ctx, v1Tov1_1Plan)

	expectedInfos := map[string]types.PlayerInfo{
		alice: {
			Index:     alice,
			LostCount: 1,
			WonCount:  2,
		},
		bob: {
			Index:     bob,
			LostCount: 1,
		},
		carol: {
			Index:     carol,
			LostCount: 1,
			WonCount:  1,
		},
	}

	for who, expectedInfo := range expectedInfos {
		storedInfo, found := suite.app.CosmoscheckersKeeper.GetPlayerInfo(suite.ctx, who)
		suite.Require().True(found)
		suite.Require().Equal(expectedInfo, storedInfo)
	}
}
