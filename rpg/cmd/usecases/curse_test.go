package usecases

import (
	"testing"

	"rpg/internal/domain"
	"rpg/internal/domain/action"
	"rpg/internal/domain/skill"
	"rpg/internal/domain/state"
)

func Test_Curse(t *testing.T) {
	rpg := domain.NewClientRPG()

	t1 := domain.NewTroop(1, rpg)
	t2 := domain.NewTroop(2, rpg)

	t1.AddRole(domain.NewRole(
		"英雄", 300, 10000, 100, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack(), skill.NewCurse()},
		action.NewHeroTxt([]string{
			"1",
			"1",
			"0",
			"0",
			"0",
			"1",
			"0",
			"1",
			"0",
		}),
	))
	t1.AddRole(domain.NewRole(
		"Ally", 600, 100, 30, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack(), skill.NewCurse(), skill.NewCurse()},
		action.NewAiI(),
	))

	t2.AddRole(domain.NewRole(
		"Slime1", 200, 999, 50, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack()},
		action.NewAiI(),
	))
	t2.AddRole(domain.NewRole(
		"Slime2", 200, 999, 100, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack()},
		action.NewAiI(),
	))

	rpg.AddTroop(t1)
	rpg.AddTroop(t2)

	rpg.Start()
}
