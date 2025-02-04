package menus

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func CreateMainMenu() tgbotapi.InlineKeyboardMarkup {
	buttons := [][]tgbotapi.InlineKeyboardButton{
		{
			tgbotapi.NewInlineKeyboardButtonData("Option 1", "option_1"),
			tgbotapi.NewInlineKeyboardButtonData("Option 2", "option_2"),
		},
		{
			tgbotapi.NewInlineKeyboardButtonData("Option 3", "option_3"),
		},
	}
	return tgbotapi.NewInlineKeyboardMarkup(buttons...)
}
