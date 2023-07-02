package state

import "rpg/internal/domain"

type petrochemicalState struct {
	state
}

func NewPetrochemicalState() domain.State {
	return &petrochemicalState{
		state{
			name:  "石化",
			round: 3,
		},
	}
}

func (s *petrochemicalState) Enter() {
	s.role.SetActionAble(false)
}

func (s *petrochemicalState) Exit() {
	s.role.SetActionAble(true)
}
