package cfsre_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/vedhavyas/cf-sre/testutil/keeper"
	"github.com/vedhavyas/cf-sre/testutil/nullify"
	"github.com/vedhavyas/cf-sre/x/cfsre"
	"github.com/vedhavyas/cf-sre/x/cfsre/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CfsreKeeper(t)
	cfsre.InitGenesis(ctx, *k, genesisState)
	got := cfsre.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
