package skill

import (
	"fmt"

	"rpg/internal/domain"
)

type curse struct {
	skill
}

func NewCurse() domain.Skill {
	return &curse{
		skill{mp: 100, name: "詛咒"},
	}
}

func (b curse) Execute(currentRole domain.Role) {
	enemies := currentRole.GetRPG().GetAllEnemies(currentRole.GetTroopID())

	output := ""
	var enemiesIndex []int
	for i, e := range enemies {
		output += fmt.Sprintf("(%d)%s ", i, e.GetName())
		enemiesIndex = append(enemiesIndex, i)
	}

	var success bool
	var targetRole domain.Role

	for !success {
		fmt.Printf(
			"選擇 1 位目標: %s\n", output,
		)

		selectedID := currentRole.ActionS2(enemiesIndex, 1)
		targetRole = enemies[selectedID[0]]
		rc := domain.NewRelationCurse(targetRole, currentRole)
		success = targetRole.AddRelationCurse(rc)
	}

	fmt.Printf("%s 對 %s 使用了 %s\n", currentRole.GetName(), targetRole.GetName(), b.name)
}
