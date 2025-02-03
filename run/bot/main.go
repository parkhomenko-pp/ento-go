package main

import (
	"ento-go/src/models"
	"errors"
	"gorm.io/driver/sqlite"
	"log"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"

	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	loadEnv()

	db := connectToDatabase()
	if db != nil {
		log.Println("Connected to database")
	}

	startBot(db)
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func connectToDatabase() *gorm.DB {
	// Подключение к базе SQLite (файл создастся автоматически, если его нет)
	db, err := gorm.Open(sqlite.Open("./tmp/data.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	// Авто-миграция (создаст таблицу, если её нет)
	if err := db.AutoMigrate(&models.Player{}); err != nil {
		log.Fatal("Ошибка миграции:", err)
	}

	return db
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

	if debug == strings.ToLower("true") {
		bot.Debug = true
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	chatString := os.Getenv("TELEGRAM_ADMIN_CHAT_ID")
	if chatString != "" {
		chatId, err := strconv.ParseInt(chatString, 10, 64)
		if err == nil {
			bot.Send(tgbotapi.NewMessage(chatId, "Bot has been started"))
		}
	}

	return bot
}

func startBot(db *gorm.DB) {
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
			var player models.Player
			result := db.First(&player, "chat_id = ?", update.Message.Chat.ID)

			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Hi, stranger"))
			} else {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Hi, "+player.Nickname))
			}

			log.Println("Got a message from: " + strconv.Itoa(int(update.Message.Chat.ID)) + "(nickname: " + update.Message.Chat.UserName + ")")
		}
	}
}
