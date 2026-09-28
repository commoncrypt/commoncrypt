package command

import (
	"context"
	"time"

	"github.com/commoncrypt/commoncrypt/internal/common/patterns"
)

type MintChallenge struct {
	Username string
}

type MintChallengeResult struct {
	ChallengeToken  string
	ChallengeSeed   string
	ChallengeExpiry time.Time
}

type MintChallengeHandler struct{}

func NewMintChallengeHandler() MintChallengeHandler {
	return MintChallengeHandler{}
}

var _ patterns.CommandProducer[MintChallenge, MintChallengeResult] = MintChallengeHandler{}

func (h MintChallengeHandler) Produce(ctx context.Context, cmd MintChallenge) (MintChallengeResult, error) {
	// TODO: implement
	return MintChallengeResult{}, nil
}
