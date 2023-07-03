package skill

import (
	"fmt"

	"rpg/internal/domain"
)

type onePunch struct {
	skill
	handler OnePunchHandler
}

func NewOnePunch(handler OnePunchHandler) domain.Skill {
	return &onePunch{
		skill:   skill{name: "一拳攻擊"},
		handler: handler,
	}
}

func (b onePunch) Execute(currentRole domain.Role) {
	enemies := currentRole.GetRPG().GetAllEnemies(currentRole.GetTroopID())

	output := ""
	var enemiesIndex []int
	for i, e := range enemies {
		output += fmt.Sprintf("(%d) %s ", i, e.GetName())
		enemiesIndex = append(enemiesIndex, i)
	}

	fmt.Printf(
		"選擇 1 位目標: %s\n", output,
	)

	selectedID := currentRole.ActionS2(enemiesIndex, 1)
	targetRole := enemies[selectedID[0]]

	fmt.Printf("%s 對 %s 使用了 %s。\n", currentRole.GetName(), targetRole.GetName(), b.name)
	b.handler.execute(currentRole, targetRole)
}
