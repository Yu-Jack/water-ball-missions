package main

import (
	"os"

	"employee.database/internal/domain"
)

func Example() {
	err := os.Setenv("PASSWORD", "1")
	if err != nil {
		panic(err) // for convenience
	}

	database := domain.NewDatabaseProxy()
	employee := database.GetEmployeeByID(5)
	employee.GetSubordinates(database)
	employee.Display()

	// Output:
	// name: peterchen, age: 3
	// name: waterball, age: 25
	// name: cc, age: 18

}
