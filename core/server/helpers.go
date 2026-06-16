package server

import "net/http"

// Chains Middleware
// handler := Chain(mux, Logger, Recover, Auth)
func Chain(h http.Handler, mw ...func(http.Handler) http.Handler) http.Handler {
	for i := len(mw) - 1; i >= 0; i-- {
		h = mw[i](h)
	}
	return h
}
