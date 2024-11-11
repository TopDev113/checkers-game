package keeper_test

import (
	"testing"

	"cosmossdk.io/collections"
	"github.com/stretchr/testify/require"
	"github.com/topdev113/checkers"
)

func TestGetGame(t *testing.T) {
	f := initFixture(t)
	require := require.New(t)

	resp, err := f.queryServer.GetGame(f.ctx, &checkers.QueryGetGameRequest{Index: "id1"})
	require.NoError(err)
	require.Equal((*checkers.StoredGame)(nil), resp.Game)
	require.Error(collections.ErrNotFound, err)

	_, err = f.msgServer.CheckersCreateGm(f.ctx, &checkers.ReqCheckersTorram{
		Creator: f.addrs[0].String(),
		Index:   "id1",
		Black:   f.addrs[1].String(),
		Red:     f.addrs[2].String(),
	})
	require.NoError(err)

	resp, err = f.queryServer.GetGame(f.ctx, &checkers.QueryGetGameRequest{Index: "id1"})
	require.NoError(err)
	require.NotEqual(int64(0), resp.Game.StartTime)
}
