package skill

import (
	"rpg/internal/domain"
)

type onePunchII struct{}

func NewOnePunchII() OnePunchHandler {
	return &onePunchII{}
}

func (o onePunchII) match(currentRole, targetRole domain.Role) bool {
	stateName := targetRole.GetState().GetName()

	if stateName == "石化" || stateName == "中毒" {
		return true
	}

	return false
}

func (o onePunchII) execute(currentRole, targetRole domain.Role) {
	damage := 80 + currentRole.GetExtraStr()

	for i := 0; i < 3; i++ {
		domain.LogDamage(currentRole.GetName(), targetRole.GetName(), damage)
		targetRole.MinusHp(damage)
	}
}
