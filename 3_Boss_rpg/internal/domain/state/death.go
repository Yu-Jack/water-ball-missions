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
	for _, rc := range s.role.GetRelationCurses() {
		rc.RecoverCurseHp()
	}

	for _, rs := range s.role.GetRelationSummons() {
		rs.RecoverSummonerHp()
	}
}
