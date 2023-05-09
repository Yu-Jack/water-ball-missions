package domain

import "time"

type Case struct {
	Symptoms     []Symptom    `json:"symptoms"`
	Prescription Prescription `json:"prescription"`
	Time         time.Time    `json:"time"`
}

func NewCase(symptoms []Symptom, prescription Prescription, time time.Time) Case {
	return Case{
		Symptoms:     symptoms,
		Prescription: prescription,
		Time:         time,
	}
}
