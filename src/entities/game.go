package entities

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

	Dots string
	Size uint8
}
