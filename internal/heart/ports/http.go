package ports

import (
	"context"

	"github.com/commoncrypt/commoncrypt/internal/heart/app"
	"github.com/commoncrypt/commoncrypt/internal/heart/app/command"
)

//go:generate go tool oapi-codegen -config oapi_codegen.json http.json

type HttpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) HttpServer {
	return HttpServer{
		app,
	}
}

var _ StrictServerInterface = HttpServer{}

func (h HttpServer) PostSendVerificationCode(
	ctx context.Context,
	request PostSendVerificationCodeRequestObject,
) (PostSendVerificationCodeResponseObject, error) {
	h.app.Commands.SendVerificationCode.Handle(ctx, command.SendVerificationCode{
		Email: *request.Body.Email,
	})

	return PostSendVerificationCode200Response{}, nil
}

func (h HttpServer) PostAccountCompleteRegistration(
	ctx context.Context,
	request PostAccountCompleteRegistrationRequestObject,
) (PostAccountCompleteRegistrationResponseObject, error) {
	return PostAccountCompleteRegistration200Response{}, nil
}
