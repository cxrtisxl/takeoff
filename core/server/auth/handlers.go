package auth

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/markbates/goth/gothic"
)

func jwksHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(cachedJWKS)
}

// OAuth (login with) endpoints

func oauthHandler(w http.ResponseWriter, r *http.Request) {
	gothUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		slog.Error(err.Error())
		gothic.BeginAuthHandler(w, r)
	} else {
		user, err := json.Marshal(gothUser)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(user)
	}
}

func oauthCallbackHandler(w http.ResponseWriter, r *http.Request) {
	gothUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}
	user, err := json.Marshal(gothUser)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(user)
}

func oauthLogoutHandler(w http.ResponseWriter, r *http.Request) {
	gothic.Logout(w, r)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}
