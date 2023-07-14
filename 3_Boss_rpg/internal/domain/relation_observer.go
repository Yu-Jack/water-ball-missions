package domain

type RelationObserver interface {
	Action()

	// GetName  attacker's name
	GetName() string
}
