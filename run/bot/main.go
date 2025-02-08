package main

import (
	"ento-go/src"
	"ento-go/src/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	entoBot := initBot()
	entoBot.Start()
}

func initBot() (entoBot *src.EntoBot) {
	// init env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	// init tg bot
	apiKey := getEnvVar("TELEGRAM_API_KEY")
	debug := getEnvVar("DEBUG")
	adminChatID, err := strconv.ParseInt(getEnvVar("TELEGRAM_ADMIN_CHAT_ID"), 10, 64)
	if err != nil {
		log.Fatalf("Invalid TELEGRAM_ADMIN_CHAT_ID: %v", err)
	}

	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = strings.ToLower(debug) == "true"
	log.Printf("Authorized on account %s", bot.Self.UserName)

	if chatString := getEnvVar("TELEGRAM_ADMIN_CHAT_ID"); chatString != "" {
		if chatId, err := strconv.ParseInt(chatString, 10, 64); err == nil {
			bot.Send(tgbotapi.NewMessage(chatId, "Bot has been started"))
		}
	}

	// init db
	db, err := gorm.Open(sqlite.Open("./tmp/data.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	if err := db.AutoMigrate(&models.Player{}); err != nil {
		log.Fatal("Ошибка миграции:", err)
	}

	entoBot = &src.EntoBot{Db: db, Tg: bot, AdminChatID: adminChatID}

	return entoBot
}

func getEnvVar(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s not set in .env file", key)
	}
	return value
}
