package testcaseas

import (
	"fmt"
	"os"

	"rpg/internal/domain/game"
)

func ExamplePetrochemical() {
	input := `#軍隊-1-開始
英雄 400 99999 30 石化
#軍隊-1-結束
#軍隊-2-開始
攻擊力超強的BOSS 270 9999 399 石化
#軍隊-2-結束
1
0
0
0
1
0
0
1
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
	//輪到 [1]英雄 (HP: 400, MP: 99999, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 石化
	//[1]英雄 對 [2]攻擊力超強的BOSS 使用了 石化。
	//輪到 [2]攻擊力超強的BOSS (HP: 270, MP: 9999, STR: 399, State: 石化)。
	//輪到 [1]英雄 (HP: 400, MP: 99899, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 石化
	//[1]英雄 攻擊 [2]攻擊力超強的BOSS。
	//[1]英雄 對 [2]攻擊力超強的BOSS 造成 30 點傷害。
	//輪到 [2]攻擊力超強的BOSS (HP: 240, MP: 9999, STR: 399, State: 石化)。
	//輪到 [1]英雄 (HP: 400, MP: 99899, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 石化
	//[1]英雄 攻擊 [2]攻擊力超強的BOSS。
	//[1]英雄 對 [2]攻擊力超強的BOSS 造成 30 點傷害。
	//輪到 [2]攻擊力超強的BOSS (HP: 210, MP: 9999, STR: 399, State: 石化)。
	//輪到 [1]英雄 (HP: 400, MP: 99899, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 石化
	//[1]英雄 攻擊 [2]攻擊力超強的BOSS。
	//[1]英雄 對 [2]攻擊力超強的BOSS 造成 30 點傷害。
	//輪到 [2]攻擊力超強的BOSS (HP: 180, MP: 9999, STR: 399, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 石化
	//[2]攻擊力超強的BOSS 攻擊 [1]英雄。
	//[2]攻擊力超強的BOSS 對 [1]英雄 造成 399 點傷害。
	//輪到 [1]英雄 (HP: 1, MP: 99899, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 石化
	//[1]英雄 對 [2]攻擊力超強的BOSS 使用了 石化。
	//輪到 [2]攻擊力超強的BOSS (HP: 180, MP: 9999, STR: 399, State: 石化)。
	//輪到 [1]英雄 (HP: 1, MP: 99799, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 石化
	//[1]英雄 攻擊 [2]攻擊力超強的BOSS。
	//[1]英雄 對 [2]攻擊力超強的BOSS 造成 30 點傷害。
	//輪到 [2]攻擊力超強的BOSS (HP: 150, MP: 9999, STR: 399, State: 石化)。
	//輪到 [1]英雄 (HP: 1, MP: 99799, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 石化
	//[1]英雄 攻擊 [2]攻擊力超強的BOSS。
	//[1]英雄 對 [2]攻擊力超強的BOSS 造成 30 點傷害。
	//輪到 [2]攻擊力超強的BOSS (HP: 120, MP: 9999, STR: 399, State: 石化)。
	//輪到 [1]英雄 (HP: 1, MP: 99799, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 石化
	//[1]英雄 對 [2]攻擊力超強的BOSS 使用了 石化。
	//輪到 [2]攻擊力超強的BOSS (HP: 120, MP: 9999, STR: 399, State: 石化)。
	//輪到 [1]英雄 (HP: 1, MP: 99699, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 石化
	//[1]英雄 攻擊 [2]攻擊力超強的BOSS。
	//[1]英雄 對 [2]攻擊力超強的BOSS 造成 30 點傷害。
	//輪到 [2]攻擊力超強的BOSS (HP: 90, MP: 9999, STR: 399, State: 石化)。
	//輪到 [1]英雄 (HP: 1, MP: 99699, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 石化
	//[1]英雄 攻擊 [2]攻擊力超強的BOSS。
	//[1]英雄 對 [2]攻擊力超強的BOSS 造成 30 點傷害。
	//輪到 [2]攻擊力超強的BOSS (HP: 60, MP: 9999, STR: 399, State: 石化)。
	//輪到 [1]英雄 (HP: 1, MP: 99699, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 石化
	//[1]英雄 攻擊 [2]攻擊力超強的BOSS。
	//[1]英雄 對 [2]攻擊力超強的BOSS 造成 30 點傷害。
	//輪到 [2]攻擊力超強的BOSS (HP: 30, MP: 9999, STR: 399, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 石化
	//[2]攻擊力超強的BOSS 對 [1]英雄 使用了 石化。
	//輪到 [1]英雄 (HP: 1, MP: 99699, STR: 30, State: 石化)。
	//輪到 [2]攻擊力超強的BOSS (HP: 30, MP: 9899, STR: 399, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 石化
	//[2]攻擊力超強的BOSS 攻擊 [1]英雄。
	//[2]攻擊力超強的BOSS 對 [1]英雄 造成 399 點傷害。
	//[1]英雄 死亡。
	//你失敗了！

}
