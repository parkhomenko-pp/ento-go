package models

import "time"

type Player struct {
	ChatID          int64 `gorm:"primaryKey;unique"`
	LastStateName   string
	LastStateValue  string
	LastStateUpdate time.Time
	Nickname        string
	ThemeId         uint8
	GamesCount      int
	WinsCount       int
}

func NewPlayer(chatID int64, nickname string) *Player {
	return &Player{
		ChatID:          chatID,
		LastStateName:   StateMainMenu,
		LastStateValue:  "",
		LastStateUpdate: time.Now(),
		Nickname:        nickname,
		ThemeId:         light,
		GamesCount:      0,
		WinsCount:       0,
	}
}

func getUserState(p *Player) PlayerState {
	return PlayerState{
		Name:       p.LastStateName,
		Value:      p.LastStateValue,
		LastActive: p.LastStateUpdate,
	}
}
