package treasure

import (
	"treasure/internal/domain"
	"treasure/internal/domain/state"
)

type dokodemoDoor struct{ treasure }

func NewDokodemoDoor(position domain.Position) domain.Treasure {
	return &dokodemoDoor{
		treasure: treasure{
			position: position,
			content:  ContentDokodemoDoor,
		},
	}
}

func (t *dokodemoDoor) GiveState() domain.State {
	return state.NewStateTeleport()
}
