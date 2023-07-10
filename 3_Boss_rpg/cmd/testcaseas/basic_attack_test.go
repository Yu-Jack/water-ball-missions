package testcaseas

import (
	"fmt"
	"os"

	"rpg/internal/domain/game"
)

func ExampleBasicAttack() {
	input := `#軍隊-1-開始
英雄 500 500 100
WaterTA 200 200 70
#軍隊-1-結束
#軍隊-2-開始
Slime1 200 90 50
Slime2 200 90 50
Slime3 200 9000 50
#軍隊-2-結束
0
0
0
0
0
1
0
0
0
0
0
`
	r, w, _ := os.Pipe()
	os.Stdin = r
	_, _ = fmt.Fprint(w, input)
	_ = w.Close()

	game.StartRPG()

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
