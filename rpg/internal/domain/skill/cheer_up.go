package skill

import (
	"fmt"

	"rpg/internal/domain"
	"rpg/internal/domain/state"
)

type cheerUp struct {
	skill
	target int
}

func NewCheerUp() domain.Skill {
	return &cheerUp{
		target: 1,
		skill:  skill{mp: 100, name: "鼓舞"},
	}
}

func (b cheerUp) Execute(currentRole domain.Role) {
	allies := currentRole.GetRPG().GetAllAllies(currentRole.GetTroopID())

	output := ""
	var indexes []int
	for i, e := range allies {
		output += fmt.Sprintf("(%d) %s ", i, e.GetName())
		indexes = append(indexes, i)
	}

	fmt.Printf(
		"選擇 %d 位目標: %s\n", b.target, output,
	)

	output = ""
	if len(indexes) > b.target {
		selectedIDs := currentRole.ActionS2(indexes, b.target)
		for _, ID := range selectedIDs {
			allies[ID].SetState(state.NewCheerUpState())
			output += fmt.Sprintf("%s ", allies[ID].GetName())
		}
	} else {
		for _, ally := range allies {
			ally.SetState(state.NewCheerUpState())
			output += fmt.Sprintf("%s ", ally.GetName())
		}
	}

	fmt.Printf("%s 對 %s 使用了 %s。\n", currentRole.GetName(), output, b.name)
}
