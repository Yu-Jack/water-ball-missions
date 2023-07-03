package skill

import (
	"rpg/internal/domain"
)

type onePunchI struct{}

func NewOnePunchI() OnePunchHandler {
	return &onePunchI{}
}

func (o onePunchI) match(currentRole, targetRole domain.Role) bool {
	if targetRole.GetHp() >= 500 {
		return true
	}
	return false
}

func (o onePunchI) execute(currentRole, targetRole domain.Role) {
	damage := 300 + currentRole.GetExtraStr()

	domain.LogDamage(currentRole.GetNameWithTroop(), targetRole.GetNameWithTroop(), damage)

	targetRole.MinusHp(damage)
}
