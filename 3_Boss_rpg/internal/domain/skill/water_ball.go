package skill

import (
	"fmt"
	"strings"

	"rpg/internal/domain"
)

type waterBall struct {
	skill
	limit int
}

func NewWaterBall() domain.Skill {
	return &waterBall{
		skill: skill{mp: 50, name: domain.SkillNameWaterBall},
		limit: 1,
	}
}

func (b waterBall) Execute(currentRole domain.Role) {
	enemies := currentRole.GetRPG().GetAllEnemies(currentRole.GetTroopID())

	var output []string
	var enemiesIndex []int
	for i, e := range enemies {
		output = append(output, fmt.Sprintf("(%d) %s", i, e.GetNameWithTroop()))
		enemiesIndex = append(enemiesIndex, i)
	}

	selectedID := currentRole.ActionS2(
		enemiesIndex, b.limit, strings.Join(output, " "),
	)
	targetRole := enemies[selectedID[0]]

	damage := 120 + currentRole.GetExtraStr()

	fmt.Printf("%s 對 %s 使用了 %s。\n", currentRole.GetNameWithTroop(), targetRole.GetNameWithTroop(), b.name)

	domain.LogDamage(currentRole.GetNameWithTroop(), targetRole.GetNameWithTroop(), damage)

	targetRole.MinusHp(damage)
}
