package usecases

import (
	"rpg/internal/domain"
	"rpg/internal/domain/action"
	"rpg/internal/domain/skill"
	"rpg/internal/domain/state"
)

func ExampleSelfExplosion() {
	rpg := domain.NewClientRPG()

	t1 := domain.NewTroop(1, rpg)
	t2 := domain.NewTroop(2, rpg)

	t1.AddRole(domain.NewRole(
		"英雄", 999999, 500, 30, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack(), skill.NewSelfExplosion()},
		action.NewHeroTxt([]string{
			"1",
		}),
	))

	for _, name := range []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"} {
		t1.AddRole(domain.NewRole(
			name, 100, 1000, 15, state.NewNormalState(),
			[]domain.Skill{skill.NewBasicAttack()},
			action.NewAiI()),
		)
	}

	for _, name := range []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"} {
		t2.AddRole(domain.NewRole(
			name, 100, 1000, 15, state.NewNormalState(),
			[]domain.Skill{skill.NewBasicAttack()},
			action.NewAiI()),
		)
	}

	rpg.AddTroop(t1)
	rpg.AddTroop(t2)

	rpg.Start()

	// Output:
	//輪到 [1]英雄 (HP: 999999, MP: 500, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 自爆
	//[1]英雄 對 [1]A, [1]B, [1]C, [1]D, [1]E, [1]F, [1]G, [1]H, [1]I, [2]A, [2]B, [2]C, [2]D, [2]E, [2]F, [2]G, [2]H, [2]I 使用了 自爆。
	//[1]英雄 對 [1]A 造成 150 點傷害。
	//[1]A 死亡。
	//[1]英雄 對 [1]B 造成 150 點傷害。
	//[1]B 死亡。
	//[1]英雄 對 [1]C 造成 150 點傷害。
	//[1]C 死亡。
	//[1]英雄 對 [1]D 造成 150 點傷害。
	//[1]D 死亡。
	//[1]英雄 對 [1]E 造成 150 點傷害。
	//[1]E 死亡。
	//[1]英雄 對 [1]F 造成 150 點傷害。
	//[1]F 死亡。
	//[1]英雄 對 [1]G 造成 150 點傷害。
	//[1]G 死亡。
	//[1]英雄 對 [1]H 造成 150 點傷害。
	//[1]H 死亡。
	//[1]英雄 對 [1]I 造成 150 點傷害。
	//[1]I 死亡。
	//[1]英雄 對 [2]A 造成 150 點傷害。
	//[2]A 死亡。
	//[1]英雄 對 [2]B 造成 150 點傷害。
	//[2]B 死亡。
	//[1]英雄 對 [2]C 造成 150 點傷害。
	//[2]C 死亡。
	//[1]英雄 對 [2]D 造成 150 點傷害。
	//[2]D 死亡。
	//[1]英雄 對 [2]E 造成 150 點傷害。
	//[2]E 死亡。
	//[1]英雄 對 [2]F 造成 150 點傷害。
	//[2]F 死亡。
	//[1]英雄 對 [2]G 造成 150 點傷害。
	//[2]G 死亡。
	//[1]英雄 對 [2]H 造成 150 點傷害。
	//[2]H 死亡。
	//[1]英雄 對 [2]I 造成 150 點傷害。
	//[2]I 死亡。
	//[1]英雄 死亡。
	//你失敗了！

}
