package app

import (
	"github.com/commoncrypt/commoncrypt/internal/heart/app/command"
	"github.com/commoncrypt/commoncrypt/internal/heart/app/query"
	"github.com/commoncrypt/commoncrypt/internal/heart/app/service"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	SendVerificationCode command.SendVerificationCodeHandler
	MintChallenge        command.MintChallengeHandler
	MintTokens           command.MintTokensHandler
}

type Queries struct {
	CheckChallengeStatus query.CheckChallengeStatusHandler
}

func NewApplication() Application {
	var (
		emailService            service.EmailService
		verificationCodeService service.VerificationCodeService
	)

	return Application{
		Commands: Commands{
			SendVerificationCode: command.NewSendVerificationCodeHandler(emailService, verificationCodeService),
			MintChallenge:        command.NewMintChallengeHandler(),
			MintTokens:           command.NewMintTokensHandler(),
		},
		Queries: Queries{
			CheckChallengeStatus: query.NewCheckChallengeStatusHandler(),
		},
	}
}
