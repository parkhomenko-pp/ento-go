package menus

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const MenuNameNewGameAcceptWaiting = "new_game_accept_waiting"

type MenuNewGameAcceptWaiting struct {
}

func (m MenuNewGameAcceptWaiting) GetName() string {
	//TODO implement me
	panic("implement me")
}

func (m MenuNewGameAcceptWaiting) DoAction() {
	//TODO implement me
	panic("implement me")
}

func (m MenuNewGameAcceptWaiting) GetFirstTimeMessage() *tgbotapi.MessageConfig {
	//TODO implement me
	panic("implement me")
}

func (m MenuNewGameAcceptWaiting) GetReplyMessage() *tgbotapi.MessageConfig {
	//TODO implement me
	panic("implement me")
}

func (m MenuNewGameAcceptWaiting) CheckReply() bool {
	//TODO implement me
	panic("implement me")
}

func (m MenuNewGameAcceptWaiting) GetOpponentMessage() *tgbotapi.MessageConfig {
	//TODO implement me
	panic("implement me")
}
