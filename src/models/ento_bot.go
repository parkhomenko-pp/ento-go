package models

import (
	"ento-go/src/models/menus"
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

const (
	StateStart        = "start"
	StateRegistration = "registration"
	StateMainMenu     = "main_menu"
)

type EntoBot struct {
	Db     *gorm.DB
	Tg     *tgbotapi.BotAPI
	States map[int64]UserState
}

func (b *EntoBot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.Tg.GetUpdatesChan(u)

	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()
	go func() {
		for range ticker.C {
			b.CleanUpInactiveStates(24 * time.Hour)
		}
	}()

	for update := range updates {
		if update.Message != nil {
			b.ReceiveMessage(update.Message)
		} else if update.CallbackQuery != nil {
			b.HandleCallbackQuery(update.CallbackQuery)
		}
	}
}

func (b *EntoBot) CleanUpInactiveStates(duration time.Duration) {
	for chatID, userState := range b.States {
		if time.Since(userState.LastActive) > duration {
			delete(b.States, chatID)
		}
	}
}

func (b *EntoBot) ReceiveMessage(message *tgbotapi.Message) {
	var player Player
	result := b.Db.First(&player, "chat_id = ?", message.Chat.ID)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		if message.Text == "/start" || message.Text == "/menu" {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Please register by providing your nickname.")
			b.Tg.Send(msg)
		} else {
			player = NewPlayer(message.Chat.ID, message.Text)
			b.Db.Create(&player)
			msg := tgbotapi.NewMessage(message.Chat.ID, "Hi, "+player.Nickname)
			msg.ReplyMarkup = menus.CreateMainMenu()
			b.Tg.Send(msg)
		}
	} else {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Hi, "+player.Nickname)
		msg.ReplyMarkup = menus.CreateMainMenu()
		b.Tg.Send(msg)
	}

	log.Println("Got a message from: " + strconv.Itoa(int(message.Chat.ID)) + " (nickname: " + message.Chat.UserName + ")")
}

func (b *EntoBot) HandleCallbackQuery(callbackQuery *tgbotapi.CallbackQuery) {
	var responseText string

	switch callbackQuery.Data {
	case "option_1":
		responseText = "You selected Option 1"
	case "option_2":
		responseText = "You selected Option 2"
	case "option_3":
		responseText = "You selected Option 3"
	default:
		responseText = "Unknown option"
	}

	msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, responseText)
	b.Tg.Send(msg)

	// Optionally, you can also answer the callback query
	callback := tgbotapi.NewCallback(callbackQuery.ID, responseText)
	b.Tg.Request(callback)
}
