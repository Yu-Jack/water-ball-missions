package skill

import (
	"fmt"
	"strings"

	"rpg/internal/domain"
)

type fireBall struct {
	skill
}

func NewFireBall() domain.Skill {
	return &fireBall{
		skill{mp: 50, name: "火球"},
	}
}

func (b fireBall) Execute(currentRole domain.Role) {
	enemies := currentRole.GetRPG().GetAllEnemies(currentRole.GetTroopID())
	damage := 50 + currentRole.GetExtraStr()

	var output []string
	for _, targetRole := range enemies {
		output = append(output, targetRole.GetName())

	}
	fmt.Printf("%s 對 %s 使用了 %s。\n", currentRole.GetName(), strings.Join(output, ", "), b.name)

	for _, targetRole := range enemies {
		domain.LogDamage(currentRole.GetName(), targetRole.GetName(), damage)
		targetRole.MinusHp(damage)
	}
}
