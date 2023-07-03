package treasure

import (
	"treasure/internal/domain"
	"treasure/internal/domain/state"
)

type poison struct{ treasure }

func NewPoison(position domain.Position) domain.Treasure {
	return &poison{
		treasure: treasure{
			position: position,
			content:  ContentPoison,
		},
	}
}

func (t *poison) GiveState() domain.State {
	return state.NewStatePoisoned()
}
