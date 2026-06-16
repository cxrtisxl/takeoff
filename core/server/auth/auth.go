package auth

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/markbates/goth"
)

var Mux = http.NewServeMux()

type WithCallback struct {
	Callback string
}

func CallbackURL(providerName string) string {
	return fmt.Sprintf("%s/auth/%s/callback", os.Getenv("URL"), providerName)
}

func Init(providers ...goth.Provider) {
	initJWT()

	// Setting up goth OAuth
	goth.UseProviders(providers...)
	var enabled string
	for k := range goth.GetProviders() {
		enabled += " " + k
	}
	slog.Info("auth providers enabled:" + enabled)

	// Creating auth handlers
	Mux.HandleFunc("GET /auth/{provider}", oauthHandler)
	Mux.HandleFunc("GET /auth/{provider}/callback", oauthCallbackHandler)
	Mux.HandleFunc("GET /logout/{provider}", oauthLogoutHandler)
	Mux.HandleFunc("GET /.well-known/jwks.json", jwksHandler)
}
