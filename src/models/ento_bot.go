package models

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	_ "time"
)

const (
	StateStart        = "start"
	StateRegistration = "registration"
	StateMainMenu     = "main_menu"
)

type EntoBot struct {
	Db      *gorm.DB
	Tg      *tgbotapi.BotAPI
	Players map[int64]*Player
}

func (b *EntoBot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.Tg.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			b.ReceiveMessage(update.Message)
		}
	}
}

func (b *EntoBot) ReceiveMessage(message *tgbotapi.Message) {
	player := b.GetPlayer(message.Chat.ID)

	if player.Nickname == "" {
		b.Tg.Send(tgbotapi.NewMessage(message.Chat.ID, "Enter your nickname. It will be used in the game."))
		player.LastStateName = StateRegistration
		b.Db.Save(&player)
	}
}

func (b *EntoBot) GetPlayer(chatID int64) *Player {
	player, ok := b.Players[chatID]
	if ok {
		return player
	}

	result := b.Db.First(&player, "chat_id = ?", chatID)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return player
	}

	player = NewPlayer(chatID, "")
	b.Db.Create(&player)
	b.Players[chatID] = player
	return player
}

func (b *EntoBot) SavePlayer(player *Player) {
	b.Db.Save(player)
	b.Players[player.ChatID] = player
}
