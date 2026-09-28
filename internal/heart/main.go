package main

import (
	"net/http"

	"github.com/commoncrypt/commoncrypt/internal/heart/app"
	"github.com/commoncrypt/commoncrypt/internal/heart/ports"
)

func main() {
	app := app.NewApplication()
	server := ports.NewHttpServer(app)
	strictHandler := ports.NewStrictHandler(server, []ports.StrictMiddlewareFunc{})
	handler := ports.Handler(strictHandler)
	http.ListenAndServe(":1984", handler)
}
