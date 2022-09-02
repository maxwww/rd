package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/maxwww/rd/bot"
	"github.com/maxwww/rd/postgres"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	token := os.Getenv("TOKEN")
	postgresURL := os.Getenv("POSTGRESQL_URL")

	botApi, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.Open(postgresURL)
	if err != nil {
		log.Fatalf("cannot open database: %v", err)
	}

	b := bot.NewBot(botApi, db)

	b.Start()
}
