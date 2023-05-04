package state

import "treasure/internal/domain"

type stockpile struct{ state }

func NewStateStockpile() domain.State {
	return &stockpile{
		state: state{
			round: 2,
			name:  TypeStockpile,
		},
	}
}

func (s *stockpile) Recover() {
	if s.round == 0 {
		newS := NewStateErupting()
		newS.SetRole(s.role)
		s.role.SetState(newS)
	}
}

func (s *stockpile) AttackedTransfer() {
	newS := NewStateNormal()
	newS.SetRole(s.role)
	s.role.SetState(newS)
}
