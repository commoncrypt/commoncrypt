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

type MintChallengeProducer struct{}

func NewMintChallengeProducer() MintChallengeProducer {
	return MintChallengeProducer{}
}

var _ patterns.CommandProducer[MintChallenge, MintChallengeResult] = MintChallengeProducer{}

func (h MintChallengeProducer) Produce(ctx context.Context, cmd MintChallenge) (MintChallengeResult, error) {
	// TODO: implement
	return MintChallengeResult{}, nil
}
