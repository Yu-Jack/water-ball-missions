//go:generate go-enum -f=$GOFILE --nocase

package domain

type State interface {
	Do()
	Enter()
	Exit()
	GetName() StateName
	CountDown()
	SetRole(r Role)
	Die()
}

// StateName
/*
ENUM(
CheerUp = "受到鼓舞"
Death = "死亡"
Normal = "正常"
Petrochemical = "石化"
Poisoned = "中毒"
)
*/
type StateName string
