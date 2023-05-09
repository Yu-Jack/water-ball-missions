package main

import "prescriber/internal/domain"

func main() {
	pf := domain.NewPrescribeFacade(
		"./database/patients.json",
		"./database/prescription.csv",
	)

	pf.Prescribe(
		"A000000000",
		[]domain.Symptom{
			domain.SymptomSnore,
		},
		func(exporter func(format, filePath string)) {
			exporter("json", "./database/result")
		},
	)
}
