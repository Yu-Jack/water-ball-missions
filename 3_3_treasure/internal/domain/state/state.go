package state

import (
	"treasure/internal/domain"
)

type state struct {
	name  string
	round int
	role  domain.Role
}

func (s *state) Recover() {
	if s.round == 0 {
		newS := NewStateNormal()
		newS.SetRole(s.role)
		s.role.SetState(newS)
	}
}

func (s *state) SetRole(role domain.Role) { s.role = role }
func (s *state) GetRound() int            { return s.round }
func (s *state) GetType() string          { return s.name }
func (s *state) IsNormal() bool           { return s.name == "Normal" }
func (s *state) IsControlledAble() bool   { return true }
func (s *state) Do()                      {}
func (s *state) Enter()                   {}
func (s *state) Exit()                    {}
func (s *state) AttackedTransfer() {
	newS := NewStateInvincible()
	newS.SetRole(s.role)
	s.role.SetState(newS)
}
func (s *state) CountDown() {
	if s.round <= 0 {
		return
	}
	s.round--
}
func (s *state) IsAttackedAble() bool { return true }
