package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/topdev113/checkers"
)

func TestInitGenesis(t *testing.T) {
	fixture := initFixture(t)

	data := &checkers.GenesisState{
		IndexedStoredGameList: []checkers.IndexedStoredGame{},
		Params:                checkers.DefaultParams(),
	}
	err := fixture.k.InitGenesis(fixture.ctx, data)
	require.NoError(t, err)

	params, err := fixture.k.Params.Get(fixture.ctx)
	require.NoError(t, err)

	require.Equal(t, checkers.DefaultParams(), params)
}

func TestExportGenesis(t *testing.T) {
	fixture := initFixture(t)

	out, err := fixture.k.ExportGenesis(fixture.ctx)
	require.NoError(t, err)

	require.Equal(t, checkers.DefaultParams(), out.Params)
	require.Equal(t, []checkers.IndexedStoredGame(nil), out.IndexedStoredGameList)
}
