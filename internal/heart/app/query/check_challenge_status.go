package query

import (
	"context"

	"github.com/commoncrypt/commoncrypt/internal/common/patterns"
	"github.com/commoncrypt/commoncrypt/internal/heart/domain/pushchallenge"
)

type CheckChallengeStatus struct {
	ChallengeToken string
}

type CheckChallengeStatusResult struct {
	Status pushchallenge.PushChallengeStatus
}

type CheckChallengeStatusProducer struct {
	repository pushchallenge.Repository
}

func NewCheckChallengeStatusProducer(repository pushchallenge.Repository) CheckChallengeStatusProducer {
	return CheckChallengeStatusProducer{
		repository: repository,
	}
}

var _ patterns.QueryProducer[CheckChallengeStatus, CheckChallengeStatusResult] = CheckChallengeStatusProducer{}

func (h CheckChallengeStatusProducer) Produce(ctx context.Context, q CheckChallengeStatus) (CheckChallengeStatusResult, error) {
	status, err := h.repository.Check(ctx, q.ChallengeToken)
	return CheckChallengeStatusResult{Status: status}, err
}
