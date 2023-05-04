//go:generate go-enum -f=$GOFILE --nocase

package domain

type State interface {
	Do()
	Enter()
	Exit()
	AttackedTransfer()
	CountDown()
	Recover()
	IsAttackedAble() bool
	GetRound() int
	SetRole(role Role)
	IsNormal() bool
	IsOrderless() bool
	GetType() string
}
