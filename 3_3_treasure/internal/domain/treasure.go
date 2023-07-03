//go:generate go-enum -f=$GOFILE --nocase

package domain

type Treasure interface {
	GetPosition() Position
	GiveState() State
	GetContent() string
}
