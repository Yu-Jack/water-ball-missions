package state

import "treasure/internal/domain"

type invincible struct{ state }

func NewStateInvincible() domain.State {
	return &invincible{
		state: state{
			round: 2,
			name:  TypeInvincible,
		},
	}
}

func (s *invincible) IsAttackedAble() bool { return false }
