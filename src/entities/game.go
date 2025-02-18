package entities

const GameStatusWaitingForAccept = 0

type Game struct {
	PlayerChatId         int64 `gorm:"primaryKey"`
	OpponentChatId       int64 `gorm:"primaryKey"`
	PlayerOpponentGameId int64 `gorm:"primaryKey"` // игр может быть несколько для одного игрока

	Status int8
}
