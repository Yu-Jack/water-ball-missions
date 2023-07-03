package state

import (
	"rpg/internal/domain"
)

type state struct {
	name  string
	round int
	role  domain.Role
}

func (s *state) Do()                   {}
func (s *state) Enter()                {}
func (s *state) Exit()                 {}
func (s *state) GetName() string       { return s.name }
func (s *state) SetRole(r domain.Role) { s.role = r }
func (s *state) Die() {
	s.role.SetState(NewDeathState())
}

func (s *state) CountDown() {
	if s.round > 0 {
		s.round--
	}
	if s.round == -1 {
		s.role.SetState(NewNormalState())
	}
}
