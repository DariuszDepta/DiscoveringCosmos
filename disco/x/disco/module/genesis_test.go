package disco_test

import (
	"testing"

	keepertest "disco/testutil/keeper"
	"disco/testutil/nullify"
	disco "disco/x/disco/module"
	"disco/x/disco/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DiscoKeeper(t)
	disco.InitGenesis(ctx, k, genesisState)
	got := disco.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
