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

	t1.CreateNewRole(
		"英雄1", 300, 300, 20, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack(), skill.NewPoison(), skill.NewFireBall(), onePunchSkill},
		action.NewHero(),
	)
	t1.CreateNewRole(
		"友人2", 300, 300, 20, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack(), skill.NewWaterBall()},
		action.NewAiI(),
	)
	t2.CreateNewRole(
		"敵人1", 400, 300, 20, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack()},
		action.NewAiI(),
	)
	t2.CreateNewRole(
		"敵人2", 300, 300, 20, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack()},
		action.NewAiI(),
	)

	rpg.AddTroop(t1)
	rpg.AddTroop(t2)

	rpg.Start()
}
