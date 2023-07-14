package state

import (
	"rpg/internal/domain"
)

type deathState struct {
	state
}

func NewDeathState() domain.State {
	return &deathState{
		state{
			name:  domain.StateNameDeath,
			round: 0,
		},
	}
}

func (s *deathState) Enter() {
	for _, rc := range s.role.GetRelations() {
		rc.Action()
	}
}
