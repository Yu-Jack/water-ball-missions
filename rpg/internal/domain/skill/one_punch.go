package skill

import (
	"fmt"
	"strings"

	"rpg/internal/domain"
)

type onePunch struct {
	skill
	handler OnePunchHandler
	limit   int
}

func NewOnePunch(handler OnePunchHandler) domain.Skill {
	return &onePunch{
		skill:   skill{name: "一拳攻擊"},
		handler: handler,
		limit:   1,
	}
}

func (b onePunch) Execute(currentRole domain.Role) {
	enemies := currentRole.GetRPG().GetAllEnemies(currentRole.GetTroopID())

	var output []string
	var enemiesIndex []int
	for i, e := range enemies {
		output = append(output, fmt.Sprintf("(%d) %s", i, e.GetName()))
		enemiesIndex = append(enemiesIndex, i)
	}

	selectedID := currentRole.ActionS2(
		enemiesIndex,
		b.limit,
		strings.Join(output, " "),
	)

	targetRole := enemies[selectedID[0]]

	fmt.Printf("%s 對 %s 使用了 %s。\n", currentRole.GetName(), targetRole.GetName(), b.name)
	b.handler.execute(currentRole, targetRole)
}
