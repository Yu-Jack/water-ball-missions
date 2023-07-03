package skill

import (
	"fmt"
	"strings"

	"rpg/internal/domain"
	"rpg/internal/domain/state"
)

type petrochemical struct {
	skill
	limit int
}

func NewPetrochemical() domain.Skill {
	return &petrochemical{
		skill: skill{mp: 100, name: "石化"},
		limit: 1,
	}
}

func (b petrochemical) Execute(currentRole domain.Role) {
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
	targetRole.SetState(state.NewPetrochemicalState())

	fmt.Printf("%s 對 %s 使用了 %s。\n", currentRole.GetNameWithTroop(), targetRole.GetNameWithTroop(), b.name)
}
