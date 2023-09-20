package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/perfogic/cosmos-checkers/x/cosmoscheckers/types"
)

var canPlayOkResponse = &types.QueryCanPlayMoveResponse{
	Possible: true,
	Reason:   "ok",
}

func (suite *IntegrationTestSuite) TestCanPlayAfterCreate() {
	suite.setupSuiteWithOneGameForPlayMove()
	goCtx := sdk.WrapSDKContext(suite.ctx)
	response, err := suite.queryClient.CanPlayMove(goCtx, &types.QueryCanPlayMoveRequest{
		GameIndex: "1",
		Player:    "b",
		FromX:     1,
		FromY:     2,
		ToX:       2,
		ToY:       3,
	})
	suite.Require().Nil(err)
	suite.Require().EqualValues(canPlayOkResponse, response)
}
