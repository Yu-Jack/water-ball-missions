package treasure

import (
	"treasure/internal/domain"
	"treasure/internal/domain/state"
)

type superStar struct{ treasure }

func NewSuperStar(position domain.Position) domain.Treasure {
	return &superStar{
		treasure: treasure{
			position: position,
			content:  ContentSuperStar,
		},
	}
}

func (t *superStar) GiveState() domain.State {
	return state.NewStateInvincible()
}
