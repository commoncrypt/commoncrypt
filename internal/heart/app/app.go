package app

import (
	"github.com/commoncrypt/commoncrypt/internal/heart/app/command"
	"github.com/commoncrypt/commoncrypt/internal/heart/app/query"
	"github.com/commoncrypt/commoncrypt/internal/heart/app/service"
	"github.com/commoncrypt/commoncrypt/internal/heart/domain/pushchallenge"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	SendVerificationCode command.SendVerificationCodeHandler
	MintChallenge        command.MintChallengeProducer
	MintTokens           command.MintTokensProducer
}

type Queries struct {
	CheckChallengeStatus query.CheckChallengeStatusProducer
}

func NewApplication() Application {
	var (
		emailService            service.EmailService
		verificationCodeService service.VerificationCodeService
		pushChallengeRepository pushchallenge.Repository
	)

	return Application{
		Commands: Commands{
			SendVerificationCode: command.NewSendVerificationCodeHandler(emailService, verificationCodeService),
			MintChallenge:        command.NewMintChallengeProducer(),
			MintTokens:           command.NewMintTokensProducer(),
		},
		Queries: Queries{
			CheckChallengeStatus: query.NewCheckChallengeStatusProducer(pushChallengeRepository),
		},
	}
}
