package skill

import (
	"fmt"
	"strings"

	"rpg/internal/domain"
	"rpg/internal/domain/state"
)

type poison struct {
	skill
}

func NewPoison() domain.Skill {
	return &poison{
		skill{mp: 80, name: "下毒"},
	}
}

func (b poison) Execute(currentRole domain.Role) {
	enemies := currentRole.GetRPG().GetAllEnemies(currentRole.GetTroopID())

	var output []string
	var enemiesIndex []int
	for i, e := range enemies {
		output = append(output, fmt.Sprintf("(%d) %s", i, e.GetNameWithTroop()))
		enemiesIndex = append(enemiesIndex, i)
	}

	selectedID := currentRole.ActionS2(
		enemiesIndex, 1, strings.Join(output, " "),
	)
	targetRole := enemies[selectedID[0]]

	fmt.Printf("%s 對 %s 使用了 %s。\n", currentRole.GetNameWithTroop(), targetRole.GetNameWithTroop(), b.name)

	targetRole.SetState(state.NewPoisonedState())
}
