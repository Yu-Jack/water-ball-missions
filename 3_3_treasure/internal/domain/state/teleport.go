package state

import "treasure/internal/domain"

type teleport struct{ state }

func NewStateTeleport() domain.State {
	return &teleport{
		state: state{
			round: 1,
			name:  "Teleport",
		},
	}
}

func (s *teleport) Exit() {
	s.role.RandomMoved()
}
