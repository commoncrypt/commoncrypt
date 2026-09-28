package command

import (
	"context"

	"github.com/commoncrypt/commoncrypt/internal/common/errors"
	"github.com/commoncrypt/commoncrypt/internal/common/patterns"
	"github.com/commoncrypt/commoncrypt/internal/heart/app/service"
	"github.com/commoncrypt/commoncrypt/internal/heart/domain/account"
)

type SendVerificationCode struct {
	Email string
}

type SendVerificationCodeHandler struct {
	emailService            service.EmailService
	verificationCodeService service.VerificationCodeService
}

func NewSendVerificationCodeHandler(
	emailService service.EmailService,
	verificationCodeService service.VerificationCodeService,
) SendVerificationCodeHandler {
	return SendVerificationCodeHandler{
		emailService,
		verificationCodeService,
	}
}

var _ patterns.CommandHandler[SendVerificationCode] = SendVerificationCodeHandler{}

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
