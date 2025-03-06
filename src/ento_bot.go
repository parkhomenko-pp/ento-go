package src

import (
	"ento-go/src/entities"
	"ento-go/src/models"
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	_ "time"
)

type EntoBot struct {
	Db *gorm.DB
	Tg *tgbotapi.BotAPI

	AdminChatID int64
}

func (b *EntoBot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.Tg.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			b.ProcessMessage(update.Message)
		}
	}
}

func (b *EntoBot) ProcessMessage(message *tgbotapi.Message) {
	// получить пользователя. если не найден, то создать нового в меню регистрации
	player := b.GetPlayer(message.Chat.ID)

	// определить меню в котором он находится
	menu := b.GetMenu(message, player)

	// сделать действие
	menu.DoAction()

	// отправить ответное сообщение текущему пользователю
	b.Tg.Send(menu.GetReplyMessage())

	// отправить ответное сообщение оппоненту
	if opponentMessage := menu.GetOpponentMessage(); opponentMessage != nil {
		b.Tg.Send(opponentMessage)
	}

	// сохранить пользователя
	b.Db.Save(&player) // TODO: только если были изменения
}

func (b *EntoBot) GetPlayer(chatID int64) *entities.Player {
	var player *entities.Player

	result := b.Db.First(&player, "chat_id = ?", chatID)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return player
	}

	player = entities.NewPlayer(chatID)
	b.Db.Create(&player)
	return player
}

func (b *EntoBot) GetMenu(message *tgbotapi.Message, player *entities.Player) *models.Menu {
	menu := new(models.Menu)
	menu.Message = message
	menu.Player = player
	menu.Db = b.Db
	menu.InitMenu()

	return menu
}
