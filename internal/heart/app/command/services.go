package command

import "context"

type EmailService interface {
	SendVerificationEmail(ctx context.Context, email string, code string) error
}

type VerificationCodeService interface {
	Make(ctx context.Context, email string) (string, error)
	Check(ctx context.Context, email string, attempt string) error
}
