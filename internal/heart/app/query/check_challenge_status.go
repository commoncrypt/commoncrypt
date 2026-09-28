package query

import (
	"context"

	"github.com/commoncrypt/commoncrypt/internal/common/patterns"
)

type ChallengeStatus int

const (
	ChallengeStatusPending ChallengeStatus = iota
	ChallengeStatusCompleted
	ChallengeStatusNotFound
)

type CheckChallengeStatus struct {
	Username       string
	ChallengeToken string
}

type CheckChallengeStatusResult struct {
	Status ChallengeStatus
}

type CheckChallengeStatusHandler struct{}

func NewCheckChallengeStatusHandler() CheckChallengeStatusHandler {
	return CheckChallengeStatusHandler{}
}

var _ patterns.QueryProducer[CheckChallengeStatus, CheckChallengeStatusResult] = CheckChallengeStatusHandler{}

func (h CheckChallengeStatusHandler) Produce(ctx context.Context, q CheckChallengeStatus) (CheckChallengeStatusResult, error) {
	// TODO: implement
	return CheckChallengeStatusResult{}, nil
}
