package models

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type EntoBot struct {
	Db *gorm.DB
	Tg *tgbotapi.BotAPI
}

func (b *EntoBot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.Tg.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			b.ReceiveMessage(update.Message)
		}
	}
}

func (b *EntoBot) ReceiveMessage(message *tgbotapi.Message) {
	var player Player
	result := b.Db.First(&player, "chat_id = ?", message.Chat.ID)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		b.Tg.Send(tgbotapi.NewMessage(message.Chat.ID, "Hi, stranger"))
	} else {
		b.Tg.Send(tgbotapi.NewMessage(message.Chat.ID, "Hi, "+player.Nickname))
	}

	log.Println("Got a message from: " + strconv.Itoa(int(message.Chat.ID)) + "(nickname: " + message.Chat.UserName + ")")
}
