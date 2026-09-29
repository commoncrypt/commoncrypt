package ports

import (
	"context"
	"errors"

	"github.com/commoncrypt/commoncrypt/internal/heart/app"
	"github.com/commoncrypt/commoncrypt/internal/heart/app/query"
	"github.com/commoncrypt/commoncrypt/internal/heart/domain/pushchallenge"
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
	res, err := h.app.Queries.CheckChallengeStatus.Produce(ctx, query.CheckChallengeStatus{ChallengeToken: *request.Body.ChallengeToken})
	if err != nil {
		return nil, err
	}

	switch res.Status {
	case pushchallenge.StatusCompleted:
		// Command: mint authentication tokens
		// Query: grab authentication tokens
		return PostPushAuthenticate200JSONResponse{Status: res.Status}, nil
	case pushchallenge.StatusNotFound:
		return PostPushAuthenticate404JSONResponse{Status: res.Status}, nil
	case pushchallenge.StatusPending:
		return PostPushAuthenticate401JSONResponse{Status: res.Status}, nil
	default:
		return nil, errors.New("unknown push challenge status")
	}
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
