package skill

import (
	"fmt"

	"rpg/internal/domain"
)

type waterBall struct {
	skill
}

func NewWaterBall() domain.Skill {
	return &waterBall{
		skill{mp: 50, name: "水球"},
	}
}

func (b waterBall) Execute(currentRole domain.Role) {
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

	damage := 50 + currentRole.GetExtraStr()

	fmt.Printf("%s 對 %s 使用了 %s。\n", currentRole.GetName(), targetRole.GetName(), b.name)

	fmt.Printf(
		"%s 對 %s 造成 %d 點傷害。\n",
		currentRole.GetName(), targetRole.GetName(), damage,
	)

	targetRole.MinusHp(damage)
}
