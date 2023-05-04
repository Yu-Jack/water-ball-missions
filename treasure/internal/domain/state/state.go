//go:generate go-enum -f=$GOFILE --nocase

package state

import (
	"treasure/internal/domain"
)

// Type
/*
ENUM(
Accelerated
Erupting
Healing
Invincible
Normal
Orderless
Poisoned
Stockpile
Teleport
)
*/
type Type string

type state struct {
	name  Type
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
func (s *state) GetType() string          { return string(s.name) }
func (s *state) IsNormal() bool           { return s.name == TypeNormal }
func (s *state) IsOrderless() bool        { return s.name == TypeOrderless }
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
