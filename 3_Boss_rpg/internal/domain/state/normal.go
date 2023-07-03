package state

import "rpg/internal/domain"

type normalState struct {
	state
}

func NewNormalState() domain.State {
	return &normalState{
		state{
			name:  "正常",
			round: 0,
		},
	}
}
