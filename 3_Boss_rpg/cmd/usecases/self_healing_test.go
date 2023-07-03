package usecases

import (
	"rpg/internal/domain"
	"rpg/internal/domain/action"
	"rpg/internal/domain/skill"
	"rpg/internal/domain/state"
)

func ExampleSelfHealing() {
	rpg := domain.NewClientRPG()

	t1 := domain.NewTroop(1, rpg)
	t2 := domain.NewTroop(2, rpg)

	t1.AddRole(domain.NewRole(
		"英雄", 500, 500, 40, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack()},
		action.NewHeroTxt([]string{
			"0",
			"0",
			"0",
			"0",
			"0",
			"0",
			"0",
			"0",
			"0",
			"0",
		}),
	))
	t2.AddRole(domain.NewRole(
		"Slime1", 100, 100, 30, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack(), skill.NewSelfHealing()},
		action.NewAiI(),
	))

	rpg.AddTroop(t1)
	rpg.AddTroop(t2)

	rpg.Start()

	// Output:
	//輪到 [1]英雄 (HP: 500, MP: 500, STR: 40, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]英雄 攻擊 [2]Slime1。
	//[1]英雄 對 [2]Slime1 造成 40 點傷害。
	//輪到 [2]Slime1 (HP: 60, MP: 100, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 自我治療
	//[2]Slime1 攻擊 [1]英雄。
	//[2]Slime1 對 [1]英雄 造成 30 點傷害。
	//輪到 [1]英雄 (HP: 470, MP: 500, STR: 40, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]英雄 攻擊 [2]Slime1。
	//[1]英雄 對 [2]Slime1 造成 40 點傷害。
	//輪到 [2]Slime1 (HP: 20, MP: 100, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 自我治療
	//[2]Slime1 使用了 自我治療。
	//輪到 [1]英雄 (HP: 470, MP: 500, STR: 40, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]英雄 攻擊 [2]Slime1。
	//[1]英雄 對 [2]Slime1 造成 40 點傷害。
	//輪到 [2]Slime1 (HP: 130, MP: 50, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 自我治療
	//[2]Slime1 攻擊 [1]英雄。
	//[2]Slime1 對 [1]英雄 造成 30 點傷害。
	//輪到 [1]英雄 (HP: 440, MP: 500, STR: 40, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]英雄 攻擊 [2]Slime1。
	//[1]英雄 對 [2]Slime1 造成 40 點傷害。
	//輪到 [2]Slime1 (HP: 90, MP: 50, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 自我治療
	//[2]Slime1 使用了 自我治療。
	//輪到 [1]英雄 (HP: 440, MP: 500, STR: 40, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]英雄 攻擊 [2]Slime1。
	//[1]英雄 對 [2]Slime1 造成 40 點傷害。
	//輪到 [2]Slime1 (HP: 200, MP: 0, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 自我治療
	//[2]Slime1 攻擊 [1]英雄。
	//[2]Slime1 對 [1]英雄 造成 30 點傷害。
	//輪到 [1]英雄 (HP: 410, MP: 500, STR: 40, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]英雄 攻擊 [2]Slime1。
	//[1]英雄 對 [2]Slime1 造成 40 點傷害。
	//輪到 [2]Slime1 (HP: 160, MP: 0, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 自我治療
	//你缺乏 MP，不能進行此行動。
	//選擇行動：(0) 普通攻擊 (1) 自我治療
	//[2]Slime1 攻擊 [1]英雄。
	//[2]Slime1 對 [1]英雄 造成 30 點傷害。
	//輪到 [1]英雄 (HP: 380, MP: 500, STR: 40, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]英雄 攻擊 [2]Slime1。
	//[1]英雄 對 [2]Slime1 造成 40 點傷害。
	//輪到 [2]Slime1 (HP: 120, MP: 0, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 自我治療
	//你缺乏 MP，不能進行此行動。
	//選擇行動：(0) 普通攻擊 (1) 自我治療
	//[2]Slime1 攻擊 [1]英雄。
	//[2]Slime1 對 [1]英雄 造成 30 點傷害。
	//輪到 [1]英雄 (HP: 350, MP: 500, STR: 40, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]英雄 攻擊 [2]Slime1。
	//[1]英雄 對 [2]Slime1 造成 40 點傷害。
	//輪到 [2]Slime1 (HP: 80, MP: 0, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 自我治療
	//你缺乏 MP，不能進行此行動。
	//選擇行動：(0) 普通攻擊 (1) 自我治療
	//[2]Slime1 攻擊 [1]英雄。
	//[2]Slime1 對 [1]英雄 造成 30 點傷害。
	//輪到 [1]英雄 (HP: 320, MP: 500, STR: 40, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]英雄 攻擊 [2]Slime1。
	//[1]英雄 對 [2]Slime1 造成 40 點傷害。
	//輪到 [2]Slime1 (HP: 40, MP: 0, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 自我治療
	//你缺乏 MP，不能進行此行動。
	//選擇行動：(0) 普通攻擊 (1) 自我治療
	//[2]Slime1 攻擊 [1]英雄。
	//[2]Slime1 對 [1]英雄 造成 30 點傷害。
	//輪到 [1]英雄 (HP: 290, MP: 500, STR: 40, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]英雄 攻擊 [2]Slime1。
	//[1]英雄 對 [2]Slime1 造成 40 點傷害。
	//[2]Slime1 死亡。
	//你獲勝了！

}
