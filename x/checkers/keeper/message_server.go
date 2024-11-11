package keeper

import (
	"context"
	"errors"
	"fmt"
	"time"

	"cosmossdk.io/collections"
	"github.com/topdev113/checkers"
	"github.com/topdev113/checkers/rules"
)

type msgServer struct {
	k Keeper
}

var _ checkers.CheckersTorramServer = msgServer{}

// NewMsgServerImpl returns an implementation of the module MsgServer interface.
func NewMsgServerImpl(keeper Keeper) checkers.CheckersTorramServer {
	return &msgServer{k: keeper}
}

// CheckersCreateGm defines the handler for the ReqCheckersTorram message.
func (ms msgServer) CheckersCreateGm(ctx context.Context, msg *checkers.ReqCheckersTorram) (*checkers.ResCheckersTorram, error) {
	if length := len([]byte(msg.Index)); checkers.MaxIndexLength < length || length < 1 {
		return nil, checkers.ErrIndexTooLong
	}
	if _, err := ms.k.StoredGames.Get(ctx, msg.Index); err == nil || errors.Is(err, collections.ErrEncoding) {
		return nil, fmt.Errorf("game already exists at index: %s", msg.Index)
	}

	newBoard := rules.New()
	storedGame := checkers.StoredGame{
		Board:     newBoard.String(),
		Turn:      rules.PieceStrings[newBoard.Turn],
		Black:     msg.Black,
		Red:       msg.Red,
		StartTime: time.Now().Unix(),
		EndTime:   0,
	}
	if err := storedGame.Validate(); err != nil {
		return nil, err
	}
	if err := ms.k.StoredGames.Set(ctx, msg.Index, storedGame); err != nil {
		return nil, err
	}

	return &checkers.ResCheckersTorram{}, nil
}
