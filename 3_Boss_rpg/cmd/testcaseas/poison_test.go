package testcaseas

import (
	"fmt"
	"os"

	"rpg/internal/domain/game"
)

func ExamplePoison() {
	input := `#軍隊-1-開始
英雄 1000 500 0 下毒
#軍隊-1-結束
#軍隊-2-開始
Slime1 120 90 50
Slime2 120 90 50
Slime3 120 9000 50
#軍隊-2-結束
1
0
1
1
1
2
1
0
1
1
1
0
`

	r, w, _ := os.Pipe()
	os.Stdin = r
	_, _ = fmt.Fprint(w, input)
	_ = w.Close()

	game.StartRPG()

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