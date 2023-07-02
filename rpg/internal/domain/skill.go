package domain

type Skill interface {
	Execute(r Role)
	GetName() string
	GetMp() int
}
