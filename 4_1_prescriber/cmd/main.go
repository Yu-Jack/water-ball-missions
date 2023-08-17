package main

import (
	"fmt"

	"prescriber/internal/domain"
)

func main() {
	pf := domain.NewPrescribeFacade(
		"./database/patients.json",
		"./database/prescription.csv",
	)

	// queue1
	pf.Prescribe(
		"A000000000",
		[]domain.Symptom{
			domain.SymptomSnore,
		},
		func(ca domain.ClientAction, c domain.Case) {
			// 1a. only print case?
			fmt.Println(c)

			// or 1b. export to database?, just whatever you want to do
			//pf.Export("json", "./database/result", c)
		},
	)

	// queue2
	pf.Prescribe(
		"A000000000",
		[]domain.Symptom{
			domain.SymptomSnore,
		},
		func(ca domain.ClientAction, c domain.Case) {
			// 1a. only print case?
			fmt.Println(c)

			// or 1b. export to database?, just whatever you want to do
			//pf.Export("json", "./database/result", c)
		},
	)

}
