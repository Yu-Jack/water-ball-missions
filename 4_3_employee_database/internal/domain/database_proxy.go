package domain

import "os"

type databaseProxy struct {
	realDB Database
}

const PASS = "1"

func NewDatabaseProxy() Database {
	return &databaseProxy{
		realDB: NewRealDatabase(),
	}
}

func (r *databaseProxy) GetEmployeeByID(ID int) Employee {
	invalid := r.isValidPassword()
	if invalid {
		panic("Wrong Password") // for convenience
	}

	return r.realDB.GetEmployeeByID(ID)
}

func (r *databaseProxy) isValidPassword() bool {
	return os.Getenv("PASSWORD") != PASS
}
