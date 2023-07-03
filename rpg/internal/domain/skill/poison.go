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
		output = append(output, fmt.Sprintf("(%d) %s", i, e.GetName()))
		enemiesIndex = append(enemiesIndex, i)
	}

	fmt.Printf(
		"選擇 1 位目標: %s\n", strings.Join(output, " "),
	)

	selectedID := currentRole.ActionS2(enemiesIndex, 1)
	targetRole := enemies[selectedID[0]]

	fmt.Printf("%s 對 %s 使用了 %s\n", currentRole.GetName(), targetRole.GetName(), b.name)

	targetRole.SetState(state.NewPoisonedState())
}
