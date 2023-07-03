package usecases

import (
	"rpg/internal/domain"
	"rpg/internal/domain/action"
	"rpg/internal/domain/skill"
	"rpg/internal/domain/state"
)

func ExampleOnePunch() {
	onePunchSkill := skill.NewOnePunch(
		skill.NewOnePunchHandler(
			skill.NewOnePunchI(), skill.NewOnePunchHandler(
				skill.NewOnePunchII(), skill.NewOnePunchHandler(
					skill.NewOnePunchIII(), skill.NewOnePunchHandler(
						skill.NewOnePunchIV(), nil,
					),
				),
			),
		),
	)

	rpg := domain.NewClientRPG()

	t1 := domain.NewTroop(1, rpg)
	t2 := domain.NewTroop(2, rpg)

	t1.AddRole(domain.NewRole(
		"英雄", 1000, 10000, 0, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack(), onePunchSkill, skill.NewPoison(), skill.NewPetrochemical(), skill.NewCheerUp()},
		action.NewHeroTxt([]string{
			"1",
			"0",
			"2",
			"1",
			"1",
			"1",
			"1",
			"0",
			"1",
			"0",
			"3",
			"0",
			"1",
			"0",
			"2",
			"1",
		}),
	))
	t2.AddRole(domain.NewRole(
		"Slime1", 601, 0, 0, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack()},
		action.NewAiI(),
	))
	t2.AddRole(domain.NewRole(
		"Slime2", 241, 0, 0, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack()},
		action.NewAiI(),
	))
	t2.AddRole(domain.NewRole(
		"Slime3", 101, 999, 0, state.NewNormalState(),
		[]domain.Skill{skill.NewBasicAttack(), onePunchSkill, onePunchSkill, skill.NewCheerUp()},
		action.NewAiI(),
	))

	rpg.AddTroop(t1)
	rpg.AddTroop(t2)

	rpg.Start()

	// Output:
	//輪到 [1]英雄 (HP: 1000, MP: 10000, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 一拳攻擊 (2) 下毒 (3) 石化 (4) 鼓舞
	//選擇 1 位目標: (0) [2]Slime1 (1) [2]Slime2 (2) [2]Slime3
	//[1]英雄 對 [2]Slime1 使用了 一拳攻擊。
	//[1]英雄 對 [2]Slime1 造成 300 點傷害。
	//輪到 [2]Slime1 (HP: 301, MP: 0, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime1 攻擊 [1]英雄。
	//[2]Slime1 對 [1]英雄 造成 0 點傷害。
	//輪到 [2]Slime2 (HP: 241, MP: 0, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime2 攻擊 [1]英雄。
	//[2]Slime2 對 [1]英雄 造成 0 點傷害。
	//輪到 [2]Slime3 (HP: 101, MP: 999, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 一拳攻擊 (2) 一拳攻擊 (3) 鼓舞
	//[2]Slime3 攻擊 [1]英雄。
	//[2]Slime3 對 [1]英雄 造成 0 點傷害。
	//輪到 [1]英雄 (HP: 1000, MP: 9820, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 一拳攻擊 (2) 下毒 (3) 石化 (4) 鼓舞
	//選擇 1 位目標: (0) [2]Slime1 (1) [2]Slime2 (2) [2]Slime3
	//[1]英雄 對 [2]Slime2 使用了 下毒。
	//輪到 [2]Slime1 (HP: 301, MP: 0, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime1 攻擊 [1]英雄。
	//[2]Slime1 對 [1]英雄 造成 0 點傷害。
	//輪到 [2]Slime2 (HP: 241, MP: 0, STR: 0, State: 中毒)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime2 攻擊 [1]英雄。
	//[2]Slime2 對 [1]英雄 造成 0 點傷害。
	//輪到 [2]Slime3 (HP: 101, MP: 999, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 一拳攻擊 (2) 一拳攻擊 (3) 鼓舞
	//[2]Slime3 對 [1]英雄 使用了 一拳攻擊。
	//[2]Slime3 對 [1]英雄 造成 300 點傷害。
	//輪到 [1]英雄 (HP: 700, MP: 9740, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 一拳攻擊 (2) 下毒 (3) 石化 (4) 鼓舞
	//選擇 1 位目標: (0) [2]Slime1 (1) [2]Slime2 (2) [2]Slime3
	//[1]英雄 對 [2]Slime2 使用了 一拳攻擊。
	//[1]英雄 對 [2]Slime2 造成 80 點傷害。
	//[1]英雄 對 [2]Slime2 造成 80 點傷害。
	//[1]英雄 對 [2]Slime2 造成 80 點傷害。
	//[2]Slime2 死亡。
	//輪到 [2]Slime1 (HP: 301, MP: 0, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime1 攻擊 [1]英雄。
	//[2]Slime1 對 [1]英雄 造成 0 點傷害。
	//輪到 [2]Slime3 (HP: 101, MP: 819, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 一拳攻擊 (2) 一拳攻擊 (3) 鼓舞
	//[2]Slime3 對 [1]英雄 使用了 一拳攻擊。
	//[2]Slime3 對 [1]英雄 造成 300 點傷害。
	//輪到 [1]英雄 (HP: 400, MP: 9560, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 一拳攻擊 (2) 下毒 (3) 石化 (4) 鼓舞
	//選擇 1 位目標: (0) [2]Slime1 (1) [2]Slime3
	//[1]英雄 對 [2]Slime1 使用了 一拳攻擊。
	//[1]英雄 對 [2]Slime1 造成 100 點傷害。
	//輪到 [2]Slime1 (HP: 201, MP: 0, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime1 攻擊 [1]英雄。
	//[2]Slime1 對 [1]英雄 造成 0 點傷害。
	//輪到 [2]Slime3 (HP: 101, MP: 639, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 一拳攻擊 (2) 一拳攻擊 (3) 鼓舞
	//[2]Slime3 對 [2]Slime1 使用了 鼓舞。
	//輪到 [1]英雄 (HP: 400, MP: 9380, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 一拳攻擊 (2) 下毒 (3) 石化 (4) 鼓舞
	//選擇 1 位目標: (0) [2]Slime1 (1) [2]Slime3
	//[1]英雄 對 [2]Slime1 使用了 一拳攻擊。
	//[1]英雄 對 [2]Slime1 造成 100 點傷害。
	//輪到 [2]Slime1 (HP: 101, MP: 0, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime1 攻擊 [1]英雄。
	//[2]Slime1 對 [1]英雄 造成 0 點傷害。
	//輪到 [2]Slime3 (HP: 101, MP: 539, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 一拳攻擊 (2) 一拳攻擊 (3) 鼓舞
	//[2]Slime3 攻擊 [1]英雄。
	//[2]Slime3 對 [1]英雄 造成 0 點傷害。
	//輪到 [1]英雄 (HP: 400, MP: 9200, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 一拳攻擊 (2) 下毒 (3) 石化 (4) 鼓舞
	//選擇 1 位目標: (0) [2]Slime1 (1) [2]Slime3
	//[1]英雄 對 [2]Slime1 使用了 石化。
	//輪到 [2]Slime1 (HP: 101, MP: 0, STR: 0, State: 石化)。
	//輪到 [2]Slime3 (HP: 101, MP: 539, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 一拳攻擊 (2) 一拳攻擊 (3) 鼓舞
	//[2]Slime3 對 [1]英雄 使用了 一拳攻擊。
	//[2]Slime3 對 [1]英雄 造成 100 點傷害。
	//輪到 [1]英雄 (HP: 300, MP: 9100, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 一拳攻擊 (2) 下毒 (3) 石化 (4) 鼓舞
	//選擇 1 位目標: (0) [2]Slime1 (1) [2]Slime3
	//[1]英雄 對 [2]Slime1 使用了 一拳攻擊。
	//[1]英雄 對 [2]Slime1 造成 80 點傷害。
	//[1]英雄 對 [2]Slime1 造成 80 點傷害。
	//[2]Slime1 死亡。
	//輪到 [2]Slime3 (HP: 101, MP: 359, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 一拳攻擊 (2) 一拳攻擊 (3) 鼓舞
	//[2]Slime3 對 [1]英雄 使用了 一拳攻擊。
	//[2]Slime3 對 [1]英雄 造成 100 點傷害。
	//輪到 [1]英雄 (HP: 200, MP: 8920, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 一拳攻擊 (2) 下毒 (3) 石化 (4) 鼓舞
	//[1]英雄 對 [2]Slime3 使用了 下毒。
	//輪到 [2]Slime3 (HP: 101, MP: 179, STR: 0, State: 中毒)。
	//選擇行動：(0) 普通攻擊 (1) 一拳攻擊 (2) 一拳攻擊 (3) 鼓舞
	//[2]Slime3 使用了 鼓舞。
	//輪到 [1]英雄 (HP: 200, MP: 8840, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 一拳攻擊 (2) 下毒 (3) 石化 (4) 鼓舞
	//[1]英雄 對 [2]Slime3 使用了 一拳攻擊。
	//[1]英雄 對 [2]Slime3 造成 80 點傷害。
	//[2]Slime3 死亡。
	//你獲勝了！

}
