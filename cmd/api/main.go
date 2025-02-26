package main

import (
	"log"
	"os"

	"github.com/Aveshek-Singha/gopherSocial/internal/db"
	"github.com/Aveshek-Singha/gopherSocial/internal/env"
	"github.com/Aveshek-Singha/gopherSocial/internal/store"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	// Retrieve the environment variable
	cfg := config{
		addr: os.Getenv("ADDR"), // Use os.Getenv instead of env.GetString
		db: dbConfig{
			addr:         os.Getenv("DB_ADDR"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
