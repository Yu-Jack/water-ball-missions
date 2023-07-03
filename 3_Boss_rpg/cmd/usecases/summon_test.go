package usecases

import (
	"rpg/internal/domain"
	"rpg/internal/domain/action"
	"rpg/internal/domain/skill"
	"rpg/internal/domain/state"
)

func ExampleSummon() {
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

	// Output:
	//輪到 [1]英雄 (HP: 500, MP: 10000, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 召喚
	//[1]英雄 使用了 召喚。
	//輪到 [1]Slime (HP: 100, MP: 0, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]Slime 攻擊 [2]Slime1。
	//[1]Slime 對 [2]Slime1 造成 50 點傷害。
	//輪到 [2]Slime1 (HP: 950, MP: 0, STR: 99, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime1 攻擊 [1]Slime。
	//[2]Slime1 對 [1]Slime 造成 99 點傷害。
	//輪到 [1]英雄 (HP: 500, MP: 9850, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 召喚
	//[1]英雄 使用了 召喚。
	//輪到 [1]Slime (HP: 1, MP: 0, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]Slime 攻擊 [2]Slime1。
	//[1]Slime 對 [2]Slime1 造成 50 點傷害。
	//輪到 [1]Slime (HP: 100, MP: 0, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]Slime 攻擊 [2]Slime1。
	//[1]Slime 對 [2]Slime1 造成 50 點傷害。
	//輪到 [2]Slime1 (HP: 850, MP: 0, STR: 99, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime1 攻擊 [1]英雄。
	//[2]Slime1 對 [1]英雄 造成 99 點傷害。
	//輪到 [1]英雄 (HP: 401, MP: 9700, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 召喚
	//[1]英雄 使用了 召喚。
	//輪到 [1]Slime (HP: 1, MP: 0, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]Slime 攻擊 [2]Slime1。
	//[1]Slime 對 [2]Slime1 造成 50 點傷害。
	//輪到 [1]Slime (HP: 100, MP: 0, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]Slime 攻擊 [2]Slime1。
	//[1]Slime 對 [2]Slime1 造成 50 點傷害。
	//輪到 [1]Slime (HP: 100, MP: 0, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]Slime 攻擊 [2]Slime1。
	//[1]Slime 對 [2]Slime1 造成 50 點傷害。
	//輪到 [2]Slime1 (HP: 700, MP: 0, STR: 99, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime1 攻擊 [1]Slime。
	//[2]Slime1 對 [1]Slime 造成 99 點傷害。
	//[1]Slime 死亡。
	//輪到 [1]英雄 (HP: 431, MP: 9550, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 召喚
	//[1]英雄 使用了 召喚。
	//輪到 [1]Slime (HP: 100, MP: 0, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]Slime 攻擊 [2]Slime1。
	//[1]Slime 對 [2]Slime1 造成 50 點傷害。
	//輪到 [1]Slime (HP: 100, MP: 0, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]Slime 攻擊 [2]Slime1。
	//[1]Slime 對 [2]Slime1 造成 50 點傷害。
	//輪到 [1]Slime (HP: 100, MP: 0, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]Slime 攻擊 [2]Slime1。
	//[1]Slime 對 [2]Slime1 造成 50 點傷害。
	//輪到 [2]Slime1 (HP: 550, MP: 0, STR: 99, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime1 攻擊 [1]Slime。
	//[2]Slime1 對 [1]Slime 造成 99 點傷害。
	//輪到 [1]英雄 (HP: 431, MP: 9400, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 召喚
	//[1]英雄 使用了 召喚。
	//輪到 [1]Slime (HP: 100, MP: 0, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]Slime 攻擊 [2]Slime1。
	//[1]Slime 對 [2]Slime1 造成 50 點傷害。
	//輪到 [1]Slime (HP: 100, MP: 0, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]Slime 攻擊 [2]Slime1。
	//[1]Slime 對 [2]Slime1 造成 50 點傷害。
	//輪到 [1]Slime (HP: 1, MP: 0, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]Slime 攻擊 [2]Slime1。
	//[1]Slime 對 [2]Slime1 造成 50 點傷害。
	//輪到 [1]Slime (HP: 100, MP: 0, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]Slime 攻擊 [2]Slime1。
	//[1]Slime 對 [2]Slime1 造成 50 點傷害。
	//輪到 [2]Slime1 (HP: 350, MP: 0, STR: 99, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime1 攻擊 [1]Slime。
	//[2]Slime1 對 [1]Slime 造成 99 點傷害。
	//輪到 [1]英雄 (HP: 431, MP: 9250, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 召喚
	//[1]英雄 使用了 召喚。
	//輪到 [1]Slime (HP: 100, MP: 0, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]Slime 攻擊 [2]Slime1。
	//[1]Slime 對 [2]Slime1 造成 50 點傷害。
	//輪到 [1]Slime (HP: 100, MP: 0, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]Slime 攻擊 [2]Slime1。
	//[1]Slime 對 [2]Slime1 造成 50 點傷害。
	//輪到 [1]Slime (HP: 1, MP: 0, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]Slime 攻擊 [2]Slime1。
	//[1]Slime 對 [2]Slime1 造成 50 點傷害。
	//輪到 [1]Slime (HP: 1, MP: 0, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]Slime 攻擊 [2]Slime1。
	//[1]Slime 對 [2]Slime1 造成 50 點傷害。
	//輪到 [1]Slime (HP: 100, MP: 0, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]Slime 攻擊 [2]Slime1。
	//[1]Slime 對 [2]Slime1 造成 50 點傷害。
	//輪到 [2]Slime1 (HP: 100, MP: 0, STR: 99, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime1 攻擊 [1]Slime。
	//[2]Slime1 對 [1]Slime 造成 99 點傷害。
	//輪到 [1]英雄 (HP: 431, MP: 9100, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 召喚
	//[1]英雄 使用了 召喚。
	//輪到 [1]Slime (HP: 100, MP: 0, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]Slime 攻擊 [2]Slime1。
	//[1]Slime 對 [2]Slime1 造成 50 點傷害。
	//輪到 [1]Slime (HP: 100, MP: 0, STR: 50, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]Slime 攻擊 [2]Slime1。
	//[1]Slime 對 [2]Slime1 造成 50 點傷害。
	//[2]Slime1 死亡。
	//你獲勝了！

}
