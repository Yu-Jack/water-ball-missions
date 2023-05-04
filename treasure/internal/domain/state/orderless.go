package state

import (
	"math/rand"

	"treasure/internal/domain"
)

type orderless struct {
	state
	options [][]domain.Direction
}

func NewStateOrderless() domain.State {
	return &orderless{
		state: state{
			round: 3,
			name:  TypeOrderless,
		},
		options: [][]domain.Direction{
			{domain.DirectionUp, domain.DirectionDown},
			{domain.DirectionLeft, domain.DirectionRight},
		},
	}
}

func (s *orderless) Enter() {
	s.role.SetActionOptions([]domain.ActionOption{domain.ActionOptionMove})
}

func (s *orderless) Do() {
	s.role.SetDirections(s.options[rand.Intn(2)])
}
func (s *orderless) Exit() {
	s.role.SetActionOptions([]domain.ActionOption{domain.ActionOptionMove, domain.ActionOptionAttack})
	s.role.SetDirections([]domain.Direction{domain.DirectionUp, domain.DirectionDown, domain.DirectionLeft, domain.DirectionRight})
}
