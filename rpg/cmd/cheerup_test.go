package main

import (
	"testing"

	"rpg/internal/domain"
	"rpg/internal/domain/action"
	"rpg/internal/domain/skill"
	"rpg/internal/domain/state"
)

func Test_CheerUp(t *testing.T) {
	rpg := domain.NewClientRPG()

	t1 := domain.NewTroop(1, rpg)
	t2 := domain.NewTroop(2, rpg)

	t1.AddRole(domain.NewRole(
		"英雄", 500, 10000, 30, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack(), skill.NewCheerUp()},
		action.NewHeroTxt([]string{
			"1",
			"0, 1, 2",
			"1",
			"2, 3, 4",
			"0",
		}),
	))
	t1.AddRole(domain.NewRole(
		"Servant1", 1000, 0, 0, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack()},
		action.NewAiI(),
	))
	t1.AddRole(domain.NewRole(
		"Servant2", 1000, 0, 0, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack()},
		action.NewAiI(),
	))
	t1.AddRole(domain.NewRole(
		"Servant3", 1000, 0, 0, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack()},
		action.NewAiI(),
	))
	t1.AddRole(domain.NewRole(
		"Servant4", 1000, 0, 0, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack()},
		action.NewAiI(),
	))
	t1.AddRole(domain.NewRole(
		"Servant5", 1000, 0, 0, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack()},
		action.NewAiI(),
	))
	t2.AddRole(domain.NewRole(
		"Slime1", 500, 0, 0, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack()},
		action.NewAiI(),
	))

	rpg.AddTroop(t1)
	rpg.AddTroop(t2)

	rpg.Start()
}
