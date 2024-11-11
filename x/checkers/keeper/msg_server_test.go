package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/topdev113/checkers"
)

func TestCreateGm(t *testing.T) {
	f := initFixture(t)
	require := require.New(t)

	testCases := []struct {
		name         string
		request      *checkers.ReqCheckersTorram
		expectErrMsg string
	}{
		{
			name: "set invalid creator (not an address)",
			request: &checkers.ReqCheckersTorram{
				Creator: "foo",
				Index:   "id1",
				Black:   f.addrs[1].String(),
				Red:     f.addrs[2].String(),
			},
			expectErrMsg: "invalid creator address",
		},
		{
			name: "create a game with valid data",
			request: &checkers.ReqCheckersTorram{
				Creator: f.addrs[0].String(),
				Index:   "id1",
				Black:   f.addrs[1].String(),
				Red:     f.addrs[2].String(),
			},
			expectErrMsg: "",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			_, err := f.msgServer.CheckersCreateGm(f.ctx, tc.request)
			if tc.expectErrMsg != "" {
				require.Error(err)
				require.ErrorContains(err, tc.expectErrMsg)
			} else {
				require.NoError(err)

				_, err := f.k.StoredGames.Get(f.ctx, tc.request.Index)
				require.NoError(err)
			}
		})
	}

}

func TestEndGm(t *testing.T) {
	f := initFixture(t)
	require := require.New(t)

	testCases := []struct {
		name         string
		request      *checkers.ReqCheckersTorramEnd
		expectErrMsg string
	}{
		{
			name: "game does not exist",
			request: &checkers.ReqCheckersTorramEnd{
				Creator: f.addrs[0].String(),
				Index:   "id2",
			},
			expectErrMsg: "game does not exist",
		},
		{
			name: "game is ended with EndTime",
			request: &checkers.ReqCheckersTorramEnd{
				Creator: f.addrs[0].String(),
				Index:   "id1",
			},
			expectErrMsg: "",
		},
	}

	// create a game first
	{
		f.msgServer.CheckersCreateGm(f.ctx, &checkers.ReqCheckersTorram{
			Creator: f.addrs[0].String(),
			Index:   "id1",
			Black:   f.addrs[1].String(),
			Red:     f.addrs[2].String(),
		})
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			_, err := f.msgServer.CheckersEndGm(f.ctx, tc.request)
			if tc.expectErrMsg != "" {
				require.Error(err)
				require.ErrorContains(err, tc.expectErrMsg)
			} else {
				require.NoError(err)

				storedGame, err := f.k.StoredGames.Get(f.ctx, tc.request.Index)
				require.NoError(err)
				require.NotEqual(int64(0), storedGame.EndTime)
			}
		})
	}

}
