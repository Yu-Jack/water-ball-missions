package domain

type RelationObserver interface {
	Action()
	GetAttackerName() string
}
