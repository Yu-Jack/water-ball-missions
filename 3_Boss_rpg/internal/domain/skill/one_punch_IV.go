package skill

import (
	"rpg/internal/domain"
)

type onePunchIV struct{}

func NewOnePunchIV() OnePunchHandler {
	return &onePunchIV{}
}

func (o onePunchIV) match(currentRole, targetRole domain.Role) bool {
	if targetRole.GetState().GetName() == domain.StateNameNormal {
		return true
	}
	return false
}

func (o onePunchIV) execute(currentRole, targetRole domain.Role) {
	damage := 100 + currentRole.GetExtraStr()

	domain.LogDamage(currentRole.GetNameWithTroop(), targetRole.GetNameWithTroop(), damage)

	targetRole.MinusHp(damage)
}
