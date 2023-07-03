package skill

import (
	"fmt"

	"rpg/internal/domain"
	"rpg/internal/domain/state"
)

type petrochemical struct {
	skill
}

func NewPetrochemical() domain.Skill {
	return &petrochemical{
		skill{mp: 100, name: "石化"},
	}
}

func (b petrochemical) Execute(currentRole domain.Role) {
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
	targetRole.SetState(state.NewPetrochemicalState())

	fmt.Printf("%s 對 %s 使用了 %s\n", currentRole.GetName(), targetRole.GetName(), b.name)
}
