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
		skill{mp: 150, name: "招喚"},
	}
}

func (b summon) Execute(currentRole domain.Role) {
	troop := currentRole.GetRPG().GetAllyTroop(currentRole.GetTroopID())
	slime := domain.NewRole(
		"Slime",
		100,
		0,
		50,
		troop.GetID(),
		state.NewNormalState(),
		[]domain.Skill{NewBasicAttack()},
		action.NewHero(),
	)

	rs := domain.NewRelationSummon(slime, currentRole)
	slime.AddRelationSummon(rs)
	troop.AddRole(slime)

	fmt.Printf("%s 招喚史萊姆\n", currentRole.GetName())
}
