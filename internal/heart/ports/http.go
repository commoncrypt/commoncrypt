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

func (h HttpServer) PostAccountRegister(
	ctx context.Context,
	request PostAccountRegisterRequestObject,
) (PostAccountRegisterResponseObject, error) {
	h.app.Commands.RegisterAccount.Handle(ctx, command.RegisterAccount{
		Email:    *request.Body.Email,
		Password: *request.Body.Password,
	})

	return PostAccountRegister200Response{}, nil
}

func (h HttpServer) PostAccountCompleteRegistration(
	ctx context.Context,
	request PostAccountCompleteRegistrationRequestObject,
) (PostAccountCompleteRegistrationResponseObject, error) {
	return PostAccountCompleteRegistration200Response{}, nil
}
