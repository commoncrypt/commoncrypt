package command

import (
	"context"
	"time"

	"github.com/commoncrypt/commoncrypt/internal/common/patterns"
)

type MintTokens struct {
	Username string
}

type MintTokensResult struct {
	AccessToken   string
	AccessExpiry  time.Time
	RefreshToken  string
	RefreshExpiry time.Time
}

type MintTokensHandler struct{}

func NewMintTokensHandler() MintTokensHandler {
	return MintTokensHandler{}
}

var _ patterns.CommandProducer[MintTokens, MintTokensResult] = MintTokensHandler{}

func (h MintTokensHandler) Produce(ctx context.Context, cmd MintTokens) (MintTokensResult, error) {
	// TODO: implement
	return MintTokensResult{}, nil
}
