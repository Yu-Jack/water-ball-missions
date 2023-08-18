package state

import "treasure/internal/domain"

type normal struct{ state }

func NewStateNormal() domain.State {
	return &normal{
		state: state{
			round: 0,
			name:  "Normal",
		},
	}
}
