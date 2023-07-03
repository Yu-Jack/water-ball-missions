package usecases

import (
	"testing"

	"rpg/internal/domain"
	"rpg/internal/domain/action"
	"rpg/internal/domain/skill"
	"rpg/internal/domain/state"
)

func Test_Summon(t *testing.T) {
	rpg := domain.NewClientRPG()

	t1 := domain.NewTroop(1, rpg)
	t2 := domain.NewTroop(2, rpg)

	t1.AddRole(domain.NewRole(
		"英雄", 500, 10000, 30, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack(), skill.NewSummon()},
		action.NewHeroTxt([]string{
			"1",
			"1",
			"1",
			"1",
			"1",
			"1",
			"1",
		}),
	))
	t2.AddRole(domain.NewRole(
		"Slime1", 1000, 0, 99, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack()},
		action.NewAiI(),
	))

	rpg.AddTroop(t1)
	rpg.AddTroop(t2)

	rpg.Start()
}
