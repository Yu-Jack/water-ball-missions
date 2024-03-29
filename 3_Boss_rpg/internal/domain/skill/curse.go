package skill

import (
	"fmt"
	"strings"

	"rpg/internal/domain"
)

type curse struct {
	skill
	limit int
}

func NewCurse() domain.Skill {
	return &curse{
		skill: skill{mp: 100, name: domain.SkillNameCurse},
		limit: 1,
	}
}

func (b curse) Execute(currentRole domain.Role) {
	enemies := currentRole.GetRPG().GetAllEnemies(currentRole.GetTroopID())

	var output []string
	var enemiesIndex []int
	for i, e := range enemies {
		output = append(output, fmt.Sprintf("(%d) %s", i, e.GetNameWithTroop()))
		enemiesIndex = append(enemiesIndex, i)
	}

	var success bool
	var targetRole domain.Role

	for !success {
		selectedID := currentRole.ActionS2(
			enemiesIndex,
			b.limit,
			strings.Join(output, " "),
		)
		targetRole = enemies[selectedID[0]]
		rc := domain.NewRelationCurse(targetRole, currentRole)
		success = targetRole.AddRelation(rc)
	}

	fmt.Printf("%s 對 %s 使用了 %s。\n", currentRole.GetNameWithTroop(), targetRole.GetNameWithTroop(), b.name)
}
