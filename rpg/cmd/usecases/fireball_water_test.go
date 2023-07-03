package usecases

import (
	"testing"

	"rpg/internal/domain"
	"rpg/internal/domain/action"
	"rpg/internal/domain/skill"
	"rpg/internal/domain/state"
)

func Test_FireWaterBall(t *testing.T) {
	rpg := domain.NewClientRPG()

	t1 := domain.NewTroop(1, rpg)
	t2 := domain.NewTroop(2, rpg)

	t1.AddRole(domain.NewRole(
		"英雄", 300, 500, 100, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack(), skill.NewFireBall(), skill.NewWaterBall()},
		action.NewHeroTxt([]string{
			"1",
			"2",
			"1",
			"2",
			"1",
			"2",
			"1",
		}),
	))
	t2.AddRole(domain.NewRole(
		"Slime1", 200, 60, 49, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack(), skill.NewFireBall()},
		action.NewAiI(),
	))
	t2.AddRole(domain.NewRole(
		"Slime2", 200, 200, 50, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack(), skill.NewFireBall(), skill.NewWaterBall()},
		action.NewAiI(),
	))

	rpg.AddTroop(t1)
	rpg.AddTroop(t2)

	rpg.Start()
}
