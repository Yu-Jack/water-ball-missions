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

	var list []string
	var enemiesIndex []int
	for i, e := range enemies {
		list = append(list, fmt.Sprintf("(%d) %s", i, e.GetName()))
		enemiesIndex = append(enemiesIndex, i)
	}

	selectedID := currentRole.ActionS2(
		enemiesIndex,
		b.limit,
		strings.Join(list, " "),
	)

	targetRole := enemies[selectedID[0]]
	damage := currentRole.GetStr() + currentRole.GetExtraStr()

	fmt.Printf("%s 攻擊 %s。\n", currentRole.GetName(), targetRole.GetName())

	domain.LogDamage(currentRole.GetName(), targetRole.GetName(), damage)
	targetRole.MinusHp(damage)
}
