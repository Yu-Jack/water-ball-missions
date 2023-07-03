package skill

import (
	"fmt"
	"strings"

	"rpg/internal/domain"
)

type basicAttack struct {
	skill
	limit int
}

func NewBasicAttack() domain.Skill {
	return &basicAttack{
		skill: skill{name: "普通攻擊"},
		limit: 1,
	}
}

func (b basicAttack) Execute(currentRole domain.Role) {
	enemies := currentRole.GetRPG().GetAllEnemies(currentRole.GetTroopID())

	var output []string
	var enemiesIndex []int
	for i, e := range enemies {
		output = append(output, fmt.Sprintf("(%d) %s", i, e.GetName()))
		enemiesIndex = append(enemiesIndex, i)
	}

	if len(enemiesIndex) != b.limit {
		fmt.Printf(
			"選擇 1 位目標: %s\n", strings.Join(output, " "),
		)
	}

	selectedID := currentRole.ActionS2(enemiesIndex, b.limit)
	targetRole := enemies[selectedID[0]]

	damage := currentRole.GetStr() + currentRole.GetExtraStr()

	fmt.Printf("%s 攻擊 %s。\n", currentRole.GetName(), targetRole.GetName())

	domain.LogDamage(currentRole.GetName(), targetRole.GetName(), damage)

	targetRole.MinusHp(damage)
}
