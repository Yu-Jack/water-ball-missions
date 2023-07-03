package skill

import (
	"fmt"

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

	fmt.Printf(
		"%s 對 %s 造成 %d 點傷害。\n",
		currentRole.GetName(), targetRole.GetName(), damage,
	)

	targetRole.MinusHp(damage)
}
