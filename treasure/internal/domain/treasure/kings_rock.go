package treasure

import (
	"treasure/internal/domain"
	"treasure/internal/domain/state"
)

type kingsRock struct{ treasure }

func NewKingsRock(position domain.Position) domain.Treasure {
	return &kingsRock{
		treasure: treasure{
			position: position,
			content:  ContentDokodemoDoor,
		},
	}
}

func (t *kingsRock) GiveState() domain.State {
	return state.NewStateStockpile()
}
