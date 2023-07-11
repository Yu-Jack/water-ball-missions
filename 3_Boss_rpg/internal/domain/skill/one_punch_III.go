package skill

import (
	"rpg/internal/domain"
	"rpg/internal/domain/state"
)

type onePunchIII struct{}

func NewOnePunchIII() OnePunchHandler {
	return &onePunchIII{}
}

func (o onePunchIII) match(currentRole, targetRole domain.Role) bool {
	if targetRole.GetState().GetName() == domain.StateNameCheerUp {
		return true
	}
	return false
}

func (o onePunchIII) execute(currentRole, targetRole domain.Role) {
	damage := 100 + currentRole.GetExtraStr()

	domain.LogDamage(currentRole.GetNameWithTroop(), targetRole.GetNameWithTroop(), damage)

	targetRole.MinusHp(damage)
	targetRole.SetState(state.NewNormalState())
}
