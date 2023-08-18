package state

import "treasure/internal/domain"

type erupting struct{ state }

func NewStateErupting() domain.State {
	return &erupting{
		state: state{
			round: 3,
			name:  "Erupting",
		},
	}
}

func (s *erupting) Recover() {
	if s.round == 0 {
		newS := NewStateTeleport()
		newS.SetRole(s.role)
		s.role.SetState(newS)
	}
}

func (s *erupting) Enter() {
	s.role.SetPower(50)
	s.role.SetAttackScope(domain.AttackScopeAll)
}
func (s *erupting) Exit() {
	// 把 Power 和 Scope 設置成原本設定
	s.role.SetAttackScope(s.role.DefaultScope())
}
