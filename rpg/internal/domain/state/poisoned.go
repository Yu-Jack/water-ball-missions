package state

import (
	"fmt"

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
	damage := 50
	fmt.Printf("%s 中毒，扣除 %d 生命\n", s.role.GetName(), damage)
	s.role.MinusHp(damage)
}
