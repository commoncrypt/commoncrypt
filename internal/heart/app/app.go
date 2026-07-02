package app

import (
	"github.com/commoncrypt/commoncrypt/internal/heart/app/command"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	SendVerificationCode command.SendVerificationCodeHandler
}

type Queries struct {
}
