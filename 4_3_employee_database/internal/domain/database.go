package domain

type Database interface {
	GetEmployeeByID(ID int) Employee
}
