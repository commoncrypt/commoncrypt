package ports

import (
	"context"

	"github.com/commoncrypt/commoncrypt/internal/heart/app"
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

func (h HttpServer) PostPushAuthenticate(
	ctx context.Context,
	request PostPushAuthenticateRequestObject,
) (PostPushAuthenticateResponseObject, error) {
	// TODO: implement
	// Query: check the challenge
	// Command: mint authentication tokens
	// Query: grab authentication tokens
	return PostPushAuthenticate200JSONResponse{}, nil
}

func (h HttpServer) PostPushChallenge(
	ctx context.Context,
	request PostPushChallengeRequestObject,
) (PostPushChallengeResponseObject, error) {
	// TODO: implement
	// Command: create the challenge
	// Query: the challenge info
	return PostPushChallenge200JSONResponse{}, nil
}
