package domain

import "fmt"

type realEmployee struct {
	name         string
	age          int
	id           int
	subordinates []Employee
}

func NewRealEmployee(name string, age, id int) Employee {
	return &realEmployee{
		name: name,
		age:  age,
		id:   id,
	}
}

func (r *realEmployee) GetSubordinates(db Database) []Employee {
	for i := 0; i < len(r.subordinates); i++ {
		r.subordinates[i] = db.GetEmployeeByID(r.subordinates[i].GetID())
	}

	return r.subordinates
}

func (r *realEmployee) GetID() int {
	return r.id
}

func (r *realEmployee) AppendSubordinates(subordinates Employee) {
	r.subordinates = append(r.subordinates, subordinates)
}

func (r *realEmployee) Display() {
	fmt.Printf("name: %s, age: %d\n", r.name, r.age)
	for _, sub := range r.subordinates {
		if sub.GetName() != "none" {
			sub.Display()
		}
	}
}

func (r *realEmployee) GetName() string {
	return r.name
}
