package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "groovy/testutil/keeper"
	"groovy/x/groovy/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.GroovyKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
