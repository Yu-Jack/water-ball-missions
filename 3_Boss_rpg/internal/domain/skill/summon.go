package skill

import (
	"fmt"

	"rpg/internal/domain"
	"rpg/internal/domain/action"
	"rpg/internal/domain/state"
)

type summon struct {
	skill
}

func NewSummon() domain.Skill {
	return &summon{
		skill{mp: 150, name: domain.SkillNameSummon},
	}
}

func (b summon) Execute(currentRole domain.Role) {
	troop := currentRole.GetRPG().GetAllyTroop(currentRole.GetTroopID())
	slime := domain.NewRole(
		"Slime",
		100,
		0,
		50,
		state.NewNormalState(),
		[]domain.Skill{NewBasicAttack()},
		action.NewAiI(),
	)

	rs := domain.NewRelationSummon(slime, currentRole)
	slime.AddRelation(rs)
	troop.AddRole(slime)

	fmt.Printf("%s 使用了 %s。\n", currentRole.GetNameWithTroop(), b.name)
}
