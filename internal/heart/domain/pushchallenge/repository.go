package pushchallenge

import (
	"context"
)

type Repository interface {
	Mint(ctx context.Context) (Model, error)
	Check(ctx context.Context, token string) (PushChallengeStatus, error)
}
