package state

import "treasure/internal/domain"

type accelerated struct{ state }

func NewStateAccelerated() domain.State {
	return &accelerated{
		state: state{
			round: 3,
			name:  "Accelerated",
		},
	}
}

func (s *accelerated) AttackedTransfer() {
	newS := NewStateNormal()
	newS.SetRole(s.role)
	s.role.SetState(newS)
}
func (s *accelerated) Enter() { s.role.SetActionCount(2) }

func (s *accelerated) Exit() {
	s.role.SetActionCount(1)
}
