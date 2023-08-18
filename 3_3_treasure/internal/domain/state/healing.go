package state

import "treasure/internal/domain"

type healing struct{ state }

func NewStateHealing() domain.State {
	return &healing{
		state: state{
			round: 5,
			name:  "Healing",
		},
	}
}

func (s *healing) Do() {
	if s.role.IsFullHP() {
		newS := NewStateNormal()
		newS.SetRole(s.role)
		s.role.SetState(newS)
		return
	}

	s.role.PlusHP(30)

	if s.role.IsFullHP() {
		newS := NewStateNormal()
		newS.SetRole(s.role)
		s.role.SetState(newS)
	}
}
