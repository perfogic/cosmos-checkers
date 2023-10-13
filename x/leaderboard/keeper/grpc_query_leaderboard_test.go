package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/perfogic/cosmos-checkers/testutil/keeper"
	"github.com/perfogic/cosmos-checkers/testutil/nullify"
	"github.com/perfogic/cosmos-checkers/x/leaderboard/types"
)

func TestLeaderboardQuery(t *testing.T) {
	keeper, ctx := keepertest.LeaderboardKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestLeaderboard(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetLeaderboardRequest
		response *types.QueryGetLeaderboardResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetLeaderboardRequest{},
			response: &types.QueryGetLeaderboardResponse{Leaderboard: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Leaderboard(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
