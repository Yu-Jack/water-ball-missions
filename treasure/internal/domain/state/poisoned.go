package state

import "treasure/internal/domain"

type poisoned struct{ state }

func NewStatePoisoned() domain.State {
	return &poisoned{
		state: state{
			round: 3,
			name:  TypePoisoned,
		},
	}
}

func (s *poisoned) Do() {
	s.role.MinusHP(30)
}
