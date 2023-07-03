package usecases

import (
	"rpg/internal/domain"
	"rpg/internal/domain/action"
	"rpg/internal/domain/skill"
	"rpg/internal/domain/state"
)

func ExampleFireWaterBall() {
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

	// Output:
	//輪到 [1]英雄 (HP: 300, MP: 500, STR: 100, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 火球 (2) 水球
	//[1]英雄 對 [2]Slime1, [2]Slime2 使用了 火球。
	//[1]英雄 對 [2]Slime1 造成 50 點傷害。
	//[1]英雄 對 [2]Slime2 造成 50 點傷害。
	//輪到 [2]Slime1 (HP: 150, MP: 60, STR: 49, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 火球
	//[2]Slime1 攻擊 [1]英雄。
	//[2]Slime1 對 [1]英雄 造成 49 點傷害。
	//輪到 [2]Slime2 (HP: 150, MP: 200, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 火球 (2) 水球
	//[2]Slime2 攻擊 [1]英雄。
	//[2]Slime2 對 [1]英雄 造成 50 點傷害。
	//輪到 [1]英雄 (HP: 201, MP: 450, STR: 100, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 火球 (2) 水球
	//選擇 1 位目標: (0) [2]Slime1 (1) [2]Slime2
	//[1]英雄 對 [2]Slime2 使用了 水球。
	//[1]英雄 對 [2]Slime2 造成 120 點傷害。
	//輪到 [2]Slime1 (HP: 150, MP: 60, STR: 49, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 火球
	//[2]Slime1 對 [1]英雄 使用了 火球。
	//[2]Slime1 對 [1]英雄 造成 50 點傷害。
	//輪到 [2]Slime2 (HP: 30, MP: 200, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 火球 (2) 水球
	//[2]Slime2 對 [1]英雄 使用了 火球。
	//[2]Slime2 對 [1]英雄 造成 50 點傷害。
	//輪到 [1]英雄 (HP: 101, MP: 400, STR: 100, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 火球 (2) 水球
	//選擇 1 位目標: (0) [2]Slime1 (1) [2]Slime2
	//[1]英雄 對 [2]Slime2 使用了 水球。
	//[1]英雄 對 [2]Slime2 造成 120 點傷害。
	//[2]Slime2 死亡。
	//輪到 [2]Slime1 (HP: 150, MP: 10, STR: 49, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 火球
	//[2]Slime1 攻擊 [1]英雄。
	//[2]Slime1 對 [1]英雄 造成 49 點傷害。
	//輪到 [1]英雄 (HP: 52, MP: 350, STR: 100, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 火球 (2) 水球
	//[1]英雄 對 [2]Slime1 使用了 水球。
	//[1]英雄 對 [2]Slime1 造成 120 點傷害。
	//輪到 [2]Slime1 (HP: 30, MP: 10, STR: 49, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 火球
	//你缺乏 MP，不能進行此行動。
	//選擇行動：(0) 普通攻擊 (1) 火球
	//[2]Slime1 攻擊 [1]英雄。
	//[2]Slime1 對 [1]英雄 造成 49 點傷害。
	//輪到 [1]英雄 (HP: 3, MP: 300, STR: 100, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 火球 (2) 水球
	//[1]英雄 對 [2]Slime1 使用了 火球。
	//[1]英雄 對 [2]Slime1 造成 50 點傷害。
	//[2]Slime1 死亡。
	//你獲勝了！

}
