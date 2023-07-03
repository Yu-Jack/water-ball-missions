package state

import (
	"rpg/internal/domain"
)

type poisonedState struct {
	state
}

func NewPoisonedState() domain.State {
	return &poisonedState{
		state{
			name:  "中毒",
			round: 3,
		},
	}
}

func (s *poisonedState) Do() {
	damage := 30
	s.role.MinusHp(damage)
}
