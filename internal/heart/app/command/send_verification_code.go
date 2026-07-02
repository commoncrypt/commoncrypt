package command

import (
	"context"

	"github.com/commoncrypt/commoncrypt/internal/common/errors"
	"github.com/commoncrypt/commoncrypt/internal/heart/domain/account"
)

type SendVerificationCode struct {
	Email string
}

type SendVerificationCodeHandler struct {
	emailService            EmailService
	verificationCodeService VerificationCodeService
}

func NewSendVerificationCodeHandler(
	emailService EmailService,
	verificationCodeService VerificationCodeService,
) SendVerificationCodeHandler {
	return SendVerificationCodeHandler{
		emailService,
		verificationCodeService,
	}
}

func (h SendVerificationCodeHandler) Handle(ctx context.Context, cmd SendVerificationCode) error {
	if !account.IsValidEmail(cmd.Email) {
		return errors.InvalidValue.WithContext("invalid email")
	}

	code, err := h.verificationCodeService.Make(ctx, cmd.Email)
	if err != nil {
		return err
	}

	err = h.emailService.SendVerificationEmail(ctx, cmd.Email, code)
	if err != nil {
		return err
	}

	return nil
}
