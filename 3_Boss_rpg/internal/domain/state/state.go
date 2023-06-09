package state

import (
	"rpg/internal/domain"
)

type state struct {
	name  domain.StateName
	round int
	role  domain.Role
}

func (s *state) Do()                       {}
func (s *state) Enter()                    {}
func (s *state) Exit()                     {}
func (s *state) GetName() domain.StateName { return s.name }
func (s *state) SetRole(r domain.Role)     { s.role = r }
func (s *state) Die() {
	s.role.SetState(NewDeathState())
}

func (s *state) CountDown() {
	if s.round > 0 {
		s.round--
	}
	if s.round == 0 {
		s.role.SetState(NewNormalState())
	}
}
