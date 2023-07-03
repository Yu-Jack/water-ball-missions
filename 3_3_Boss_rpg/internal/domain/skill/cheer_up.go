package skill

import (
	"fmt"
	"strings"

	"rpg/internal/domain"
	"rpg/internal/domain/state"
)

type cheerUp struct {
	skill
	limit int
}

func NewCheerUp() domain.Skill {
	return &cheerUp{
		limit: 3,
		skill: skill{mp: 100, name: "鼓舞"},
	}
}

func (b cheerUp) Execute(currentRole domain.Role) {
	allies := currentRole.GetRPG().GetAllAlliesExcludeSelf(currentRole.GetTroopID(), currentRole.GetID())

	var output []string
	var list []string
	var indexes []int
	for i, e := range allies {
		list = append(list, fmt.Sprintf("(%d) %s", i, e.GetNameWithTroop()))
		indexes = append(indexes, i)
	}

	if len(indexes) == 0 {
		fmt.Printf("%s 使用了 %s。\n", currentRole.GetNameWithTroop(), b.name)
		return
	}

	if len(indexes) > b.limit {
		selectedIDs := currentRole.ActionS2(
			indexes,
			b.limit,
			strings.Join(list, " "),
		)
		for _, ID := range selectedIDs {
			allies[ID].SetState(state.NewCheerUpState())
			output = append(output, fmt.Sprintf("%s", allies[ID].GetNameWithTroop()))
		}
	} else {
		for _, ally := range allies {
			ally.SetState(state.NewCheerUpState())
			output = append(output, fmt.Sprintf("%s", ally.GetNameWithTroop()))
		}
	}

	fmt.Printf("%s 對 %s 使用了 %s。\n", currentRole.GetNameWithTroop(), strings.Join(output, ", "), b.name)
}
