package domain

type State interface {
	Do()
	Enter()
	Exit()
	GetName() string
	CountDown()
	SetRole(r Role)
	Die()
}
