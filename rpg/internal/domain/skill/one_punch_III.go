package skill

import (
	"fmt"

	"rpg/internal/domain"
	"rpg/internal/domain/state"
)

type onePunchIII struct{}

func NewOnePunchIII() OnePunchHandler {
	return &onePunchIII{}
}

func (o onePunchIII) match(currentRole, targetRole domain.Role) bool {
	if targetRole.GetState().GetName() == "鼓舞" {
		return true
	}
	return false
}

func (o onePunchIII) execute(currentRole, targetRole domain.Role) {
	damage := 100 + currentRole.GetExtraStr()

	fmt.Printf(
		"%s 對 %s 造成 %d 點傷害。\n",
		currentRole.GetName(), targetRole.GetName(), damage,
	)

	targetRole.MinusHp(damage)
	targetRole.SetState(state.NewNormalState())
}
