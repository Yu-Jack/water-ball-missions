package usecases

import (
	"rpg/internal/domain"
	"rpg/internal/domain/action"
	"rpg/internal/domain/skill"
	"rpg/internal/domain/state"
)

func ExamplePoison() {
	rpg := domain.NewClientRPG()

	t1 := domain.NewTroop(1, rpg)
	t2 := domain.NewTroop(2, rpg)

	t1.AddRole(domain.NewRole(
		"英雄", 1000, 500, 0, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack(), skill.NewPoison()},
		action.NewHeroTxt([]string{
			"1",
			"0",
			"1",
			"1",
			"1",
			"2",
			"1",
			"0",
			"1",
			"1",
			"1",
			"0",
		}),
	))
	t2.AddRole(domain.NewRole(
		"Slime1", 120, 90, 50, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack()},
		action.NewAiI(),
	))
	t2.AddRole(domain.NewRole(
		"Slime2", 120, 90, 50, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack()},
		action.NewAiI(),
	))
	t2.AddRole(domain.NewRole(
		"Slime3", 120, 9000, 50, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack()},
		action.NewAiI(),
	))

	rpg.AddTroop(t1)
	rpg.AddTroop(t2)

	rpg.Start()

	// Output:
	//輪到 [1]英雄 (HP: 1000, MP: 500, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 下毒
	//選擇 1 位目標: (0) [2]Slime1 (1) [2]Slime2 (2) [2]Slime3
	//[1]英雄 對 [2]Slime1 使用了 下毒。
	//輪到 [2]Slime1 (HP: 120, MP: 90, STR: 50, State: 中毒)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime1 攻擊 [1]英雄。
	//[2]Slime1 對 [1]英雄 造成 50 點傷害。
	//輪到 [2]Slime2 (HP: 120, MP: 90, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime2 攻擊 [1]英雄。
	//[2]Slime2 對 [1]英雄 造成 50 點傷害。
	//輪到 [2]Slime3 (HP: 120, MP: 9000, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime3 攻擊 [1]英雄。
	//[2]Slime3 對 [1]英雄 造成 50 點傷害。
	//輪到 [1]英雄 (HP: 850, MP: 420, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 下毒
	//選擇 1 位目標: (0) [2]Slime1 (1) [2]Slime2 (2) [2]Slime3
	//[1]英雄 對 [2]Slime2 使用了 下毒。
	//輪到 [2]Slime1 (HP: 90, MP: 90, STR: 50, State: 中毒)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime1 攻擊 [1]英雄。
	//[2]Slime1 對 [1]英雄 造成 50 點傷害。
	//輪到 [2]Slime2 (HP: 120, MP: 90, STR: 50, State: 中毒)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime2 攻擊 [1]英雄。
	//[2]Slime2 對 [1]英雄 造成 50 點傷害。
	//輪到 [2]Slime3 (HP: 120, MP: 9000, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime3 攻擊 [1]英雄。
	//[2]Slime3 對 [1]英雄 造成 50 點傷害。
	//輪到 [1]英雄 (HP: 700, MP: 340, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 下毒
	//選擇 1 位目標: (0) [2]Slime1 (1) [2]Slime2 (2) [2]Slime3
	//[1]英雄 對 [2]Slime3 使用了 下毒。
	//輪到 [2]Slime1 (HP: 60, MP: 90, STR: 50, State: 中毒)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime1 攻擊 [1]英雄。
	//[2]Slime1 對 [1]英雄 造成 50 點傷害。
	//輪到 [2]Slime2 (HP: 90, MP: 90, STR: 50, State: 中毒)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime2 攻擊 [1]英雄。
	//[2]Slime2 對 [1]英雄 造成 50 點傷害。
	//輪到 [2]Slime3 (HP: 120, MP: 9000, STR: 50, State: 中毒)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime3 攻擊 [1]英雄。
	//[2]Slime3 對 [1]英雄 造成 50 點傷害。
	//輪到 [1]英雄 (HP: 550, MP: 260, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 下毒
	//選擇 1 位目標: (0) [2]Slime1 (1) [2]Slime2 (2) [2]Slime3
	//[1]英雄 對 [2]Slime1 使用了 下毒。
	//輪到 [2]Slime1 (HP: 30, MP: 90, STR: 50, State: 中毒)。
	//[2]Slime1 死亡。
	//輪到 [2]Slime2 (HP: 60, MP: 90, STR: 50, State: 中毒)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime2 攻擊 [1]英雄。
	//[2]Slime2 對 [1]英雄 造成 50 點傷害。
	//輪到 [2]Slime3 (HP: 90, MP: 9000, STR: 50, State: 中毒)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime3 攻擊 [1]英雄。
	//[2]Slime3 對 [1]英雄 造成 50 點傷害。
	//輪到 [1]英雄 (HP: 450, MP: 180, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 下毒
	//選擇 1 位目標: (0) [2]Slime2 (1) [2]Slime3
	//[1]英雄 對 [2]Slime3 使用了 下毒。
	//輪到 [2]Slime2 (HP: 30, MP: 90, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime2 攻擊 [1]英雄。
	//[2]Slime2 對 [1]英雄 造成 50 點傷害。
	//輪到 [2]Slime3 (HP: 60, MP: 9000, STR: 50, State: 中毒)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime3 攻擊 [1]英雄。
	//[2]Slime3 對 [1]英雄 造成 50 點傷害。
	//輪到 [1]英雄 (HP: 350, MP: 100, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 下毒
	//選擇 1 位目標: (0) [2]Slime2 (1) [2]Slime3
	//[1]英雄 對 [2]Slime2 使用了 下毒。
	//輪到 [2]Slime2 (HP: 30, MP: 90, STR: 50, State: 中毒)。
	//[2]Slime2 死亡。
	//輪到 [2]Slime3 (HP: 30, MP: 9000, STR: 50, State: 中毒)。
	//[2]Slime3 死亡。
	//你獲勝了！

}
