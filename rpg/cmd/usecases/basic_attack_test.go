package usecases

import (
	"rpg/internal/domain"
	"rpg/internal/domain/action"
	"rpg/internal/domain/skill"
	"rpg/internal/domain/state"
)

func ExampleBasicAttack() {
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

	// Output:
	//輪到 [1]英雄 (HP: 500, MP: 500, STR: 100, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//選擇 1 位目標: (0) [2]Slime1 (1) [2]Slime2 (2) [2]Slime3
	//[1]英雄 攻擊 [2]Slime1。
	//[1]英雄 對 [2]Slime1 造成 100 點傷害。
	//輪到 [1]WaterTA (HP: 200, MP: 200, STR: 70, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]WaterTA 攻擊 [2]Slime2。
	//[1]WaterTA 對 [2]Slime2 造成 70 點傷害。
	//輪到 [2]Slime1 (HP: 100, MP: 90, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime1 攻擊 [1]WaterTA。
	//[2]Slime1 對 [1]WaterTA 造成 50 點傷害。
	//輪到 [2]Slime2 (HP: 130, MP: 90, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime2 攻擊 [1]WaterTA。
	//[2]Slime2 對 [1]WaterTA 造成 50 點傷害。
	//輪到 [2]Slime3 (HP: 200, MP: 9000, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime3 攻擊 [1]WaterTA。
	//[2]Slime3 對 [1]WaterTA 造成 50 點傷害。
	//輪到 [1]英雄 (HP: 500, MP: 500, STR: 100, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//選擇 1 位目標: (0) [2]Slime1 (1) [2]Slime2 (2) [2]Slime3
	//[1]英雄 攻擊 [2]Slime1。
	//[1]英雄 對 [2]Slime1 造成 100 點傷害。
	//[2]Slime1 死亡。
	//輪到 [1]WaterTA (HP: 50, MP: 200, STR: 70, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]WaterTA 攻擊 [2]Slime3。
	//[1]WaterTA 對 [2]Slime3 造成 70 點傷害。
	//輪到 [2]Slime2 (HP: 130, MP: 90, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime2 攻擊 [1]WaterTA。
	//[2]Slime2 對 [1]WaterTA 造成 50 點傷害。
	//[1]WaterTA 死亡。
	//輪到 [2]Slime3 (HP: 130, MP: 9000, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime3 攻擊 [1]英雄。
	//[2]Slime3 對 [1]英雄 造成 50 點傷害。
	//輪到 [1]英雄 (HP: 450, MP: 500, STR: 100, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//選擇 1 位目標: (0) [2]Slime2 (1) [2]Slime3
	//[1]英雄 攻擊 [2]Slime3。
	//[1]英雄 對 [2]Slime3 造成 100 點傷害。
	//輪到 [2]Slime2 (HP: 130, MP: 90, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime2 攻擊 [1]英雄。
	//[2]Slime2 對 [1]英雄 造成 50 點傷害。
	//輪到 [2]Slime3 (HP: 30, MP: 9000, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime3 攻擊 [1]英雄。
	//[2]Slime3 對 [1]英雄 造成 50 點傷害。
	//輪到 [1]英雄 (HP: 350, MP: 500, STR: 100, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//選擇 1 位目標: (0) [2]Slime2 (1) [2]Slime3
	//[1]英雄 攻擊 [2]Slime2。
	//[1]英雄 對 [2]Slime2 造成 100 點傷害。
	//輪到 [2]Slime2 (HP: 30, MP: 90, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime2 攻擊 [1]英雄。
	//[2]Slime2 對 [1]英雄 造成 50 點傷害。
	//輪到 [2]Slime3 (HP: 30, MP: 9000, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime3 攻擊 [1]英雄。
	//[2]Slime3 對 [1]英雄 造成 50 點傷害。
	//輪到 [1]英雄 (HP: 250, MP: 500, STR: 100, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//選擇 1 位目標: (0) [2]Slime2 (1) [2]Slime3
	//[1]英雄 攻擊 [2]Slime2。
	//[1]英雄 對 [2]Slime2 造成 100 點傷害。
	//[2]Slime2 死亡。
	//輪到 [2]Slime3 (HP: 30, MP: 9000, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime3 攻擊 [1]英雄。
	//[2]Slime3 對 [1]英雄 造成 50 點傷害。
	//輪到 [1]英雄 (HP: 200, MP: 500, STR: 100, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]英雄 攻擊 [2]Slime3。
	//[1]英雄 對 [2]Slime3 造成 100 點傷害。
	//[2]Slime3 死亡。
	//你獲勝了！
}
