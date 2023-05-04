package treasure

import (
	"treasure/internal/domain"
	"treasure/internal/domain/state"
)

type healingPotion struct{ treasure }

func NewHealingPotion(position domain.Position) domain.Treasure {
	return &healingPotion{
		treasure: treasure{
			position: position,
			content:  ContentHealingPotion,
		},
	}
}

func (t *healingPotion) GiveState() domain.State {
	return state.NewStateHealing()
}
