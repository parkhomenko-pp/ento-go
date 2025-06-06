package menus

import (
	"ento-go/src/entities"
	"testing"
)

func TestIsMyTurn(t *testing.T) {
	tests := []struct {
		playerChatID uint64
		player       entities.Player
		opponent     entities.Player
		isPlayerTurn bool
		expected     bool
	}{
		{playerChatID: 1, player: entities.Player{ChatID: 1}, opponent: entities.Player{ChatID: 2}, isPlayerTurn: false, expected: false},
		{playerChatID: 1, player: entities.Player{ChatID: 1}, opponent: entities.Player{ChatID: 2}, isPlayerTurn: true, expected: true},
		{playerChatID: 2, player: entities.Player{ChatID: 1}, opponent: entities.Player{ChatID: 2}, isPlayerTurn: false, expected: true},
		{playerChatID: 2, player: entities.Player{ChatID: 1}, opponent: entities.Player{ChatID: 2}, isPlayerTurn: true, expected: false},
	}

	for i, test := range tests {
		m := &MenuGame{
			Game: &entities.Game{
				PlayerChatID: int64(test.playerChatID),
				Player:       test.player,
				Opponent:     test.opponent,
				IsPlayerTurn: test.isPlayerTurn,
			},
		}

		if m.isMyTurn() != test.expected {
			t.Errorf("test %d: expected isMyTurn = %v, got %v", i, test.expected, m.isMyTurn())
		}
	}
}
