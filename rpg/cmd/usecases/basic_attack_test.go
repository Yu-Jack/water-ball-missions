package usecases

import (
	"testing"

	"rpg/internal/domain"
	"rpg/internal/domain/action"
	"rpg/internal/domain/skill"
	"rpg/internal/domain/state"
)

func Test_BasicAttack(t *testing.T) {
	rpg := domain.NewClientRPG()

	t1 := domain.NewTroop(1, rpg)
	t2 := domain.NewTroop(2, rpg)

	t1.AddRole(domain.NewRole(
		"英雄", 500, 500, 100, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack()},
		action.NewHeroTxt([]string{
			"0",
			"0",
			"0",
			"0",
			"0",
			"1",
			"0",
			"0",
			"0",
			"0",
			"0",
		}),
	))
	t1.AddRole(domain.NewRole(
		"WaterTA", 200, 200, 70, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack()},
		action.NewAiI(),
	))
	t2.AddRole(domain.NewRole(
		"Slime1", 200, 90, 50, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack()},
		action.NewAiI(),
	))
	t2.AddRole(domain.NewRole(
		"Slime2", 200, 90, 50, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack()},
		action.NewAiI(),
	))
	t2.AddRole(domain.NewRole(
		"Slime3", 200, 9000, 50, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack()},
		action.NewAiI(),
	))

	rpg.AddTroop(t1)
	rpg.AddTroop(t2)

	rpg.Start()
}
