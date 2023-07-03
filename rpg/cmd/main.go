package main

import (
	"rpg/internal/domain"
	"rpg/internal/domain/action"
	"rpg/internal/domain/skill"
	"rpg/internal/domain/state"
)

func main() {

	rpg := domain.NewClientRPG()

	t1 := domain.NewTroop(1, rpg)
	t2 := domain.NewTroop(2, rpg)

	onePunchSkill := skill.NewOnePunch(
		skill.NewOnePunchHandler(
			skill.NewOnePunchI(), skill.NewOnePunchHandler(
				skill.NewOnePunchII(), skill.NewOnePunchHandler(
					skill.NewOnePunchIII(), skill.NewOnePunchHandler(
						skill.NewOnePunchIV(), nil,
					),
				),
			),
		),
	)
	_ = onePunchSkill

	t1.AddRole(domain.NewRole(
		"英雄", 500, 500, 40, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack()},
		action.NewHero(),
	))
	t2.AddRole(domain.NewRole(
		"Slime1", 100, 100, 30, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack(), skill.NewSelfHealing()},
		action.NewAiI(),
	))

	rpg.AddTroop(t1)
	rpg.AddTroop(t2)

	rpg.Start()
}
