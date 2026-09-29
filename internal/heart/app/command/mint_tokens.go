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

type MintTokensProducer struct{}

func NewMintTokensProducer() MintTokensProducer {
	return MintTokensProducer{}
}

var _ patterns.CommandProducer[MintTokens, MintTokensResult] = MintTokensProducer{}

func (h MintTokensProducer) Produce(ctx context.Context, cmd MintTokens) (MintTokensResult, error) {
	// TODO: implement
	return MintTokensResult{}, nil
}
