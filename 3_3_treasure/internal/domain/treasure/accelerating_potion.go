package treasure

import (
	"treasure/internal/domain"
	"treasure/internal/domain/state"
)

type acceleratingPotion struct{ treasure }

func NewAcceleratingPotion(position domain.Position) domain.Treasure {
	return &acceleratingPotion{
		treasure: treasure{
			position: position,
			content:  ContentAcceleratingPotion,
		},
	}
}

func (t *acceleratingPotion) GiveState() domain.State {
	return state.NewStateAccelerated()
}
