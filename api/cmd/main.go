package main

import (
	"log"

	"github.com/joho/godotenv"
	"ivpn.net/email/api/config"
	"ivpn.net/email/api/internal/cron"
	"ivpn.net/email/api/internal/repository"
	"ivpn.net/email/api/internal/service"
	"ivpn.net/email/api/internal/transport/api"
)

func Run() error {
	godotenv.Load(".env") // ignore error — Zeabur injects env vars directly
	cfg, err := config.New()
	if err != nil {
		return err
	}

	db, err := repository.NewDB(cfg.DB)
	if err != nil {
		return err
	}

	cache := repository.NewMemCache()

	cron.New(db.Client)

	service := service.New(cfg, db, cache)

	err = api.Start(cfg.API, service, cache)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := Run()
	if err != nil {
		log.Fatal(err)
	}
}
