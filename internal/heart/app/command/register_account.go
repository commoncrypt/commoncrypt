package command

import (
	"context"

	"github.com/commoncrypt/commoncrypt/internal/common/errors"
	"github.com/commoncrypt/commoncrypt/internal/heart/domain/account"
)

type RegisterAccount struct {
	Email    string
	Password string
}

type RegisterAccountHandler struct {
	repo                    account.Repository
	emailService            EmailService
	verificationCodeService VerificationCodeService
}

func NewRegisterAccountHandler(
	repo account.Repository,
	emailService EmailService,
	verificationCodeService VerificationCodeService,
) RegisterAccountHandler {
	return RegisterAccountHandler{
		repo,
		emailService,
		verificationCodeService,
	}
}

func (h RegisterAccountHandler) Handle(ctx context.Context, cmd RegisterAccount) error {
	if !account.IsValidEmail(cmd.Email) {
		return errors.InvalidValue.WithContext("invalid email")
	}

	exists, err := h.repo.AccountExists(ctx, cmd.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.AlreadyExists
	}

	code, err := h.verificationCodeService.Make(ctx, cmd.Email)
	if err != nil {
		return nil
	}

	err = h.emailService.SendVerificationEmail(ctx, cmd.Email, code)
	if err != nil {
		return nil
	}

	return nil
}
