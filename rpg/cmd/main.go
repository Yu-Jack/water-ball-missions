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

	t1.AddRole(domain.NewRole(
		"英雄1", 300, 300, 20, t1.GetID(), state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack(), skill.NewCurse(), skill.NewSummon(), onePunchSkill},
		action.NewHero(),
	))
	t1.AddRole(domain.NewRole(
		"友人2", 300, 300, 20, t1.GetID(), state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack(), skill.NewWaterBall(), onePunchSkill},
		action.NewHero(),
	))
	t2.AddRole(domain.NewRole(
		"敵人1", 400, 300, 100, t1.GetID(), state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack()},
		action.NewHero(),
	))
	t2.AddRole(domain.NewRole(
		"敵人2", 100, 300, 100, t1.GetID(), state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack()},
		action.NewHero(),
	))

	rpg.AddTroop(t1)
	rpg.AddTroop(t2)

	rpg.Start()
}
