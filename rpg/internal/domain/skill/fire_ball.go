package skill

import (
	"fmt"

	"rpg/internal/domain"
)

type fireBall struct {
	skill
}

func NewFireBall() domain.Skill {
	return &fireBall{
		skill{mp: 250, name: "火球"},
	}
}

func (b fireBall) Execute(currentRole domain.Role) {
	enemies := currentRole.GetRPG().GetAllEnemies(currentRole.GetTroopID())
	damage := 50 + currentRole.GetExtraStr()

	output1 := ""
	for _, targetRole := range enemies {
		output1 += targetRole.GetName() + ", "

	}
	fmt.Printf("%s 對 %s 使用了 %s\n", currentRole.GetName(), output1, b.name)

	for _, targetRole := range enemies {
		fmt.Printf(
			"%s 對 %s 造成 %d 傷害。\n",
			currentRole.GetName(), targetRole.GetName(), damage,
		)
		targetRole.MinusHp(damage)
	}
}
