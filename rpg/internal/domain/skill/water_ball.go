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
		skill: skill{mp: 50, name: "水球"},
		limit: 1,
	}
}

func (b waterBall) Execute(currentRole domain.Role) {
	enemies := currentRole.GetRPG().GetAllEnemies(currentRole.GetTroopID())

	var output []string
	var enemiesIndex []int
	for i, e := range enemies {
		output = append(output, fmt.Sprintf("(%d) %s", i, e.GetName()))
		enemiesIndex = append(enemiesIndex, i)
	}

	fmt.Printf(
		"選擇 %d 位目標: %s\n", b.limit, strings.Join(output, " "),
	)

	selectedID := currentRole.ActionS2(enemiesIndex, b.limit)
	targetRole := enemies[selectedID[0]]

	damage := 50 + currentRole.GetExtraStr()

	fmt.Printf("%s 對 %s 使用了 %s。\n", currentRole.GetName(), targetRole.GetName(), b.name)

	domain.LogDamage(currentRole.GetName(), targetRole.GetName(), damage)

	targetRole.MinusHp(damage)
}
