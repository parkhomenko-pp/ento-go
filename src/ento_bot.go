package src

import (
	"ento-go/src/models"
	"ento-go/src/models/menus"
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gorm.io/gorm"
	"log"
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
	if message.Chat.ID != b.AdminChatID { // TODO: remove after release
		b.Tg.Send(tgbotapi.NewMessage(message.Chat.ID, "Sorry, but I can't talk with you 😔\nDeveloper is working on me"))
		return
	}

	// получить пользователя
	player := b.GetPlayer(message.Chat.ID)

	// определить меню в котором он находится
	menu := b.GetMenu(message, player)

	// сделать действие
	menu.DoAction()

	// ответить на сообщение
	b.Tg.Send(menu.GetReplyMessage())

	log.Println(menu)
}

func (b *EntoBot) GetPlayer(chatID int64) *models.Player {
	var player *models.Player

	result := b.Db.First(&player, "chat_id = ?", chatID)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return player
	}

	player = models.NewPlayer(chatID)
	b.Db.Create(&player)
	return player
}

func (b *EntoBot) GetMenu(message *tgbotapi.Message, player *models.Player) *models.Menu {
	menu := new(models.Menu)
	menu.Message = message
	menu.Player = player

	switch player.LastMenu {
	case menus.MenuRegistration:
		menu.Menuable = &menus.Registration{}
	case menus.MenuMain:
		menu.Menuable = &menus.Main{}

	default:
		if player.Nickname == "" {
			menu.Menuable = &menus.Registration{}
		} else {
			menu.Menuable = &menus.NotFound{}
		}
	}

	return menu
}
