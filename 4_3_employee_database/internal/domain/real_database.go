package domain

import (
	"strconv"
	"strings"
)

type realDatabase struct {
	data map[int]Employee
}

func NewRealDatabase() Database {
	return &realDatabase{
		data: make(map[int]Employee),
	}
}

// Simulate the CSV format, but we don't focus on too much implementation here.
const fakeData = `id name age subordinateIds
1 waterball 25
2 fixiabis 15 1,3
3 fong 7 1
4 cc 18 1,2,3
5 peterchen 3 1,4
6 handsomeboy 22 1
`

func (r *realDatabase) GetEmployeeByID(ID int) Employee {
	splitFakeData := strings.Split(fakeData, "\n")
	if ID >= len(splitFakeData) {
		panic("Failed") // for convenience
	}

	data := strings.Split(splitFakeData[ID], " ")

	name := data[1]
	age, _ := strconv.Atoi(data[2])
	employee := NewRealEmployee(name, age, ID)

	if len(data) > 3 {
		for _, id := range strings.Split(data[3], ",") {
			intID, _ := strconv.Atoi(id)
			e := NewRealEmployee("none", -1, intID)
			employee.AppendSubordinates(e)
		}
	}

	return employee
}
