package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "disco/testutil/keeper"
	"disco/x/disco/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.DiscoKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
