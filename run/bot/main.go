package main

import (
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
	db, tgBot := initParams()
	entoBot := models.EntoBot{Db: db, Tg: tgBot}
	entoBot.Start()
}

func initParams() (*gorm.DB, *tgbotapi.BotAPI) {
	// init env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	// init tg bot
	apiKey := getEnvVar("TELEGRAM_API_KEY")
	debug := getEnvVar("DEBUG")

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

	return db, bot
}

func getEnvVar(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%s not set in .env file", key)
	}
	return value
}
