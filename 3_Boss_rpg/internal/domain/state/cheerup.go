package state

import "rpg/internal/domain"

type cheerUpState struct {
	state
}

func NewCheerUpState() domain.State {
	return &cheerUpState{
		state{
			name:  domain.StateNameCheerUp,
			round: 3,
		},
	}
}

func (s *cheerUpState) Enter() {
	s.role.SetExtraStr(50)
}

func (s *cheerUpState) Exit() {
	s.role.SetExtraStr(0)
}
