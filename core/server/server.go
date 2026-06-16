package server

import (
	"log/slog"
	"net/http"

	"github.com/cxrtisxl/takeoff/core/server/auth"
	"github.com/cxrtisxl/takeoff/core/server/handler"
	mw "github.com/cxrtisxl/takeoff/core/server/middleware"
	"github.com/markbates/goth"
)

var mux = http.NewServeMux()

func Run(authProviders []goth.Provider) {
	// Add new routes here
	mux.HandleFunc("GET /{$}", handler.Root)

	// Auth
	if len(authProviders) > 0 {
		mux.Handle("/auth/", auth.Mux)
		mux.Handle("/logout/", auth.Mux)
		mux.Handle("/.well-known/", auth.Mux)
		auth.Init(authProviders...)
	} else {
		slog.Info("goth auth disabled")
	}

	http.ListenAndServe(
		":8080",
		Chain(mux, mw.Log),
	)
}
