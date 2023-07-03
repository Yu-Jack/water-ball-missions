package usecases

import (
	"rpg/internal/domain"
	"rpg/internal/domain/action"
	"rpg/internal/domain/skill"
	"rpg/internal/domain/state"
)

func ExampleCurse() {
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

	// Output:
	//輪到 [1]英雄 (HP: 300, MP: 10000, STR: 100, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 詛咒
	//選擇 1 位目標: (0) [2]Slime1 (1) [2]Slime2
	//[1]英雄 對 [2]Slime2 使用了 詛咒。
	//輪到 [1]Ally (HP: 600, MP: 100, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 詛咒 (2) 詛咒
	//[1]Ally 攻擊 [2]Slime2。
	//[1]Ally 對 [2]Slime2 造成 30 點傷害。
	//輪到 [2]Slime1 (HP: 200, MP: 999, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime1 攻擊 [1]Ally。
	//[2]Slime1 對 [1]Ally 造成 50 點傷害。
	//輪到 [2]Slime2 (HP: 170, MP: 999, STR: 100, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime2 攻擊 [1]Ally。
	//[2]Slime2 對 [1]Ally 造成 100 點傷害。
	//輪到 [1]英雄 (HP: 300, MP: 9900, STR: 100, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 詛咒
	//選擇 1 位目標: (0) [2]Slime1 (1) [2]Slime2
	//[1]英雄 攻擊 [2]Slime1。
	//[1]英雄 對 [2]Slime1 造成 100 點傷害。
	//輪到 [1]Ally (HP: 450, MP: 100, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 詛咒 (2) 詛咒
	//[1]Ally 對 [2]Slime2 使用了 詛咒。
	//輪到 [2]Slime1 (HP: 100, MP: 999, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime1 攻擊 [1]Ally。
	//[2]Slime1 對 [1]Ally 造成 50 點傷害。
	//輪到 [2]Slime2 (HP: 170, MP: 999, STR: 100, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime2 攻擊 [1]Ally。
	//[2]Slime2 對 [1]Ally 造成 100 點傷害。
	//輪到 [1]英雄 (HP: 300, MP: 9900, STR: 100, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 詛咒
	//選擇 1 位目標: (0) [2]Slime1 (1) [2]Slime2
	//[1]英雄 攻擊 [2]Slime2。
	//[1]英雄 對 [2]Slime2 造成 100 點傷害。
	//輪到 [1]Ally (HP: 300, MP: 0, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 詛咒 (2) 詛咒
	//你缺乏 MP，不能進行此行動。
	//選擇行動：(0) 普通攻擊 (1) 詛咒 (2) 詛咒
	//你缺乏 MP，不能進行此行動。
	//選擇行動：(0) 普通攻擊 (1) 詛咒 (2) 詛咒
	//[1]Ally 攻擊 [2]Slime2。
	//[1]Ally 對 [2]Slime2 造成 30 點傷害。
	//輪到 [2]Slime1 (HP: 100, MP: 999, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime1 攻擊 [1]Ally。
	//[2]Slime1 對 [1]Ally 造成 50 點傷害。
	//輪到 [2]Slime2 (HP: 40, MP: 999, STR: 100, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime2 攻擊 [1]Ally。
	//[2]Slime2 對 [1]Ally 造成 100 點傷害。
	//輪到 [1]英雄 (HP: 300, MP: 9900, STR: 100, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 詛咒
	//選擇 1 位目標: (0) [2]Slime1 (1) [2]Slime2
	//[1]英雄 攻擊 [2]Slime2。
	//[1]英雄 對 [2]Slime2 造成 100 點傷害。
	//[2]Slime2 死亡。
	//輪到 [1]Ally (HP: 1149, MP: 0, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 詛咒 (2) 詛咒
	//你缺乏 MP，不能進行此行動。
	//選擇行動：(0) 普通攻擊 (1) 詛咒 (2) 詛咒
	//[1]Ally 攻擊 [2]Slime1。
	//[1]Ally 對 [2]Slime1 造成 30 點傷害。
	//輪到 [2]Slime1 (HP: 70, MP: 999, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime1 攻擊 [1]Ally。
	//[2]Slime1 對 [1]Ally 造成 50 點傷害。
	//輪到 [1]英雄 (HP: 1299, MP: 9900, STR: 100, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 詛咒
	//[1]英雄 攻擊 [2]Slime1。
	//[1]英雄 對 [2]Slime1 造成 100 點傷害。
	//[2]Slime1 死亡。
	//你獲勝了！
}
