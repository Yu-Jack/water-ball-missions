package treasure

import (
	"treasure/internal/domain"
	"treasure/internal/domain/state"
)

type devilFruit struct{ treasure }

func NewDevilFruit(position domain.Position) domain.Treasure {
	return &devilFruit{
		treasure: treasure{
			position: position,
			content:  ContentDevilFruit,
		},
	}
}

func (t *devilFruit) GiveState() domain.State {
	return state.NewStateOrderless()
}
