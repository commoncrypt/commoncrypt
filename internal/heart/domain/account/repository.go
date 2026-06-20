package account

import "context"

type Repository interface {
	AccountExists(ctx context.Context, email string) (bool, error)
	CreateAccount(ctx context.Context, email, password string) error
}
