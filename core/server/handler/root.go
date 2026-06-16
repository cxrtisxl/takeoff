package handler

import (
	"log"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	log.Printf("ROOT path=%q", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"Time to Takeoff!"}`))
}
