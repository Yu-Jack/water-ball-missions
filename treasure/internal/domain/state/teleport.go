package state

import "treasure/internal/domain"

type teleport struct{ state }

func NewStateTeleport() domain.State {
	return &teleport{
		state: state{
			round: 1,
			name:  TypeTeleport,
		},
	}
}

func (s *teleport) Exit() {
	s.role.RandomMoved()
}
