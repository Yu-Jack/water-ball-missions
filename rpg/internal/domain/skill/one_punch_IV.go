package skill

import (
	"fmt"

	"rpg/internal/domain"
)

type onePunchIV struct{}

func NewOnePunchIV() OnePunchHandler {
	return &onePunchIV{}
}

func (o onePunchIV) match(currentRole, targetRole domain.Role) bool {
	if targetRole.GetState().GetName() == "正常" {
		return true
	}
	return false
}

func (o onePunchIV) execute(currentRole, targetRole domain.Role) {
	damage := 100 + currentRole.GetExtraStr()

	fmt.Printf(
		"%s 對 %s 造成 %d 點傷害。\n",
		currentRole.GetName(), targetRole.GetName(), damage,
	)

	targetRole.MinusHp(damage)
}
