package domain

type Employee interface {
	GetSubordinates(db Database) []Employee
	GetID() int
	AppendSubordinates(subordinates Employee)
	Display()
	GetName() string
}
