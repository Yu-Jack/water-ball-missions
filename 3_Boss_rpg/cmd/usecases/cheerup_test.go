package usecases

import (
	"fmt"
	"os"

	"rpg/internal/domain/game"
)

func ExampleCheerUp() {
	input := `#軍隊-1-開始
英雄 500 10000 30 鼓舞
Servant1 1000 0 0
Servant2 1000 0 0
Servant3 1000 0 0
Servant4 1000 0 0
Servant5 1000 0 0
#軍隊-1-結束
#軍隊-2-開始
Slime1 500 0 0
#軍隊-2-結束
1
0, 1, 2
1
2, 3, 4
0
`
	r, w, _ := os.Pipe()
	os.Stdin = r
	_, _ = fmt.Fprint(w, input)
	_ = w.Close()

	game.SetupRPGGame()

	// Output:
	//輪到 [1]英雄 (HP: 500, MP: 10000, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 鼓舞
	//選擇 3 位目標: (0) [1]Servant1 (1) [1]Servant2 (2) [1]Servant3 (3) [1]Servant4 (4) [1]Servant5
	//[1]英雄 對 [1]Servant1, [1]Servant2, [1]Servant3 使用了 鼓舞。
	//輪到 [1]Servant1 (HP: 1000, MP: 0, STR: 0, State: 受到鼓舞)。
	//選擇行動：(0) 普通攻擊
	//[1]Servant1 攻擊 [2]Slime1。
	//[1]Servant1 對 [2]Slime1 造成 50 點傷害。
	//輪到 [1]Servant2 (HP: 1000, MP: 0, STR: 0, State: 受到鼓舞)。
	//選擇行動：(0) 普通攻擊
	//[1]Servant2 攻擊 [2]Slime1。
	//[1]Servant2 對 [2]Slime1 造成 50 點傷害。
	//輪到 [1]Servant3 (HP: 1000, MP: 0, STR: 0, State: 受到鼓舞)。
	//選擇行動：(0) 普通攻擊
	//[1]Servant3 攻擊 [2]Slime1。
	//[1]Servant3 對 [2]Slime1 造成 50 點傷害。
	//輪到 [1]Servant4 (HP: 1000, MP: 0, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]Servant4 攻擊 [2]Slime1。
	//[1]Servant4 對 [2]Slime1 造成 0 點傷害。
	//輪到 [1]Servant5 (HP: 1000, MP: 0, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[1]Servant5 攻擊 [2]Slime1。
	//[1]Servant5 對 [2]Slime1 造成 0 點傷害。
	//輪到 [2]Slime1 (HP: 350, MP: 0, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime1 攻擊 [1]Servant1。
	//[2]Slime1 對 [1]Servant1 造成 0 點傷害。
	//輪到 [1]英雄 (HP: 500, MP: 9900, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 鼓舞
	//選擇 3 位目標: (0) [1]Servant1 (1) [1]Servant2 (2) [1]Servant3 (3) [1]Servant4 (4) [1]Servant5
	//[1]英雄 對 [1]Servant3, [1]Servant4, [1]Servant5 使用了 鼓舞。
	//輪到 [1]Servant1 (HP: 1000, MP: 0, STR: 0, State: 受到鼓舞)。
	//選擇行動：(0) 普通攻擊
	//[1]Servant1 攻擊 [2]Slime1。
	//[1]Servant1 對 [2]Slime1 造成 50 點傷害。
	//輪到 [1]Servant2 (HP: 1000, MP: 0, STR: 0, State: 受到鼓舞)。
	//選擇行動：(0) 普通攻擊
	//[1]Servant2 攻擊 [2]Slime1。
	//[1]Servant2 對 [2]Slime1 造成 50 點傷害。
	//輪到 [1]Servant3 (HP: 1000, MP: 0, STR: 0, State: 受到鼓舞)。
	//選擇行動：(0) 普通攻擊
	//[1]Servant3 攻擊 [2]Slime1。
	//[1]Servant3 對 [2]Slime1 造成 50 點傷害。
	//輪到 [1]Servant4 (HP: 1000, MP: 0, STR: 0, State: 受到鼓舞)。
	//選擇行動：(0) 普通攻擊
	//[1]Servant4 攻擊 [2]Slime1。
	//[1]Servant4 對 [2]Slime1 造成 50 點傷害。
	//輪到 [1]Servant5 (HP: 1000, MP: 0, STR: 0, State: 受到鼓舞)。
	//選擇行動：(0) 普通攻擊
	//[1]Servant5 攻擊 [2]Slime1。
	//[1]Servant5 對 [2]Slime1 造成 50 點傷害。
	//輪到 [2]Slime1 (HP: 100, MP: 0, STR: 0, State: 正常)。
	//選擇行動：(0) 普通攻擊
	//[2]Slime1 攻擊 [1]Servant3。
	//[2]Slime1 對 [1]Servant3 造成 0 點傷害。
	//輪到 [1]英雄 (HP: 500, MP: 9800, STR: 30, State: 正常)。
	//選擇行動：(0) 普通攻擊 (1) 鼓舞
	//[1]英雄 攻擊 [2]Slime1。
	//[1]英雄 對 [2]Slime1 造成 30 點傷害。
	//輪到 [1]Servant1 (HP: 1000, MP: 0, STR: 0, State: 受到鼓舞)。
	//選擇行動：(0) 普通攻擊
	//[1]Servant1 攻擊 [2]Slime1。
	//[1]Servant1 對 [2]Slime1 造成 50 點傷害。
	//輪到 [1]Servant2 (HP: 1000, MP: 0, STR: 0, State: 受到鼓舞)。
	//選擇行動：(0) 普通攻擊
	//[1]Servant2 攻擊 [2]Slime1。
	//[1]Servant2 對 [2]Slime1 造成 50 點傷害。
	//[2]Slime1 死亡。
	//你獲勝了！

}
