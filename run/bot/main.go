package main

import (
	"database/sql"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	loadEnv()

	db := connectToDatabase()
	defer db.Close()

	startBot()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func connectToDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./db/data.db")
	if err != nil {
		println("data-notfound")
		copyEmptyDatabase()
		db, err = sql.Open("sqlite3", "./db/data.db")
		if err != nil {
			log.Fatalf("Error opening data.db file: %v", err)
		}
	}
	return db
}

func copyEmptyDatabase() {
	input, err := os.ReadFile("./db/empty.db")
	if err != nil {
		log.Fatalf("Error reading empty.db file: %v", err)
	}

	err = os.WriteFile("./db/data.db", input, 0644)
	if err != nil {
		log.Fatalf("Error writing data.db file: %v", err)
	}

	log.Println("Successfully copied empty.db to data.db")
}

func initializeBot(apiKey string) *tgbotapi.BotAPI {
	debug := os.Getenv("DEBUG")
	if debug == "" {
		log.Fatalf("DEBUG not set in .env file")
	}

	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		log.Panic(err)
	}

	if debug == "TRUE" {
		bot.Debug = true
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return bot
}

func startBot() {
	apiKey := os.Getenv("TELEGRAM_API_KEY")
	if apiKey == "" {
		log.Fatalf("TELEGRAM_API_KEY not set in .env file")
	}

	bot := initializeBot(apiKey)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
