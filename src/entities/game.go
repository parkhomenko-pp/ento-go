package entities

import "math"

const GameStatusWaitingForAccept = 0
const GameStatusPlaying = 1
const GameStatusFinished = 2
const GameStatusDeclined = 4

type Game struct {
	ID             uint `gorm:"primaryKey"`
	Status         int8
	PlayerChatID   int64
	Player         Player `gorm:"foreignKey:PlayerChatID;references:ChatID"`
	OpponentChatID int64
	Opponent       Player `gorm:"foreignKey:OpponentChatID;references:ChatID"`
	IsPlayerTurn   bool

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
	size := int(math.Sqrt(float64(len(g.Dots))))
	board := make([][]uint8, size)

	for i := range board {
		board[i] = make([]uint8, size)
		for j := range board[i] {
			board[i][j] = g.Dots[i*size+j] - '0'
		}
	}

	return board
}
