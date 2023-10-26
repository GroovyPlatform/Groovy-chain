package groovy_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "groovy/testutil/keeper"
	"groovy/testutil/nullify"
	"groovy/x/groovy"
	"groovy/x/groovy/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GroovyKeeper(t)
	groovy.InitGenesis(ctx, *k, genesisState)
	got := groovy.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
