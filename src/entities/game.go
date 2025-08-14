package entities

import (
	"encoding/json"
)

const GameStatusWaitingForAccept = 0
const GameStatusPlaying = 1
const GameStatusFinished = 2
const GameStatusDeclined = 4

type Game struct {
	ID                uint `gorm:"primaryKey"`
	Status            int8
	PlayerChatID      int64
	Player            Player `gorm:"foreignKey:PlayerChatID;references:ChatID"`
	OpponentChatID    int64
	Opponent          Player `gorm:"foreignKey:OpponentChatID;references:ChatID"`
	IsPlayerBlack     bool   `gorm:"default:true"`
	IsPlayerTurn      bool
	LastStonePosition string // JSON string of the last stone position

	Dots string
	Size uint8
}

func (g *Game) GetOpponentChatIdForPlayer(player *Player) Player {
	if player.ChatID == g.PlayerChatID {
		return g.Opponent
	}
	return g.Player
}

func (g *Game) GetDots() [][]uint8 {
	var matrix [][]uint8
	err := json.Unmarshal([]byte(g.Dots), &matrix)
	if err != nil {
		return make([][]uint8, 0)
	}
	return matrix
}

func (g *Game) SetDots(dots [][]uint8) error {
	data, err := json.Marshal(dots)
	if err != nil {
		return err
	}
	g.Dots = string(data)
	return nil
}

func (g *Game) ToggleIsPlayerTurn() {
	g.IsPlayerTurn = !g.IsPlayerTurn
}
