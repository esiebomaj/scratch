package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppEnvs struct {
	DB_URL          string
	SCRAPE_INTERVAL int
	PORT            string
}

func ConfigureEnvs() (AppEnvs, error) {
	godotenv.Load()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("PORT not configured")
	}

	DB_URL := os.Getenv("DB_URL")
	if DB_URL == "" {
		log.Fatal("DB_URL not configured")
	}

	SCRAPE_INTERVAL := os.Getenv("SCRAPE_INTERVAL")
	if SCRAPE_INTERVAL == "" {
		log.Fatal("SCRAPE_INTERVAL not configured")
	}

	ScrapeInterval, err := strconv.Atoi(SCRAPE_INTERVAL)
	if err != nil {
		log.Fatal("Error converting string to int:", err)
	}

	return AppEnvs{
		DB_URL:          DB_URL,
		SCRAPE_INTERVAL: ScrapeInterval,
		PORT:            PORT,
	}, nil
}
