package main

import (
	"log"
	"os"

	"github.com/cxrtisxl/takeoff/core/server"
	"github.com/cxrtisxl/takeoff/core/server/auth"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/twitterv2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	server.Run(
		[]goth.Provider{
			twitterv2.NewAuthenticate(os.Getenv("X_KEY"), os.Getenv("X_SECRET"), auth.CallbackURL("twitterv2")),
		},
	)
}

// TODO
// github.com/jackc/pgx/v5            // Postgres (pgxpool)
