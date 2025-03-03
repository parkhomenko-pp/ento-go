package menus

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const MenuNameGame = "game"

type MenuGame struct {
}

func (m MenuGame) GetName() string {
	return MenuNameGame
}

func (m MenuGame) DoAction() {
	//TODO implement me
	panic("implement me")
}

func (m MenuGame) GetReplyMessage() *tgbotapi.MessageConfig {
	//TODO implement me
	panic("implement me")
}

func (m MenuGame) CheckReply() bool {
	//TODO implement me
	panic("implement me")
}

func (m MenuGame) GetOpponentMessage() *tgbotapi.MessageConfig {
	//TODO implement me
	panic("implement me")
}
