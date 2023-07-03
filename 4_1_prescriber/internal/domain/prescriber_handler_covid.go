package domain

import "time"

type prescriberHandlerCOVID struct{}

func NewPrescriberHandlerCOVID() PrescriberHandler {
	return &prescriberHandlerCOVID{}
}

func (ph *prescriberHandlerCOVID) Match(_ Patient, symptoms []Symptom) bool {
	count := 0

	for _, symptom := range symptoms {
		if symptom == SymptomCough {
			count++
		}
		if symptom == SymptomHeadache {
			count++
		}
		if symptom == SymptomSneeze {
			count++
		}
	}

	return count == 3 && len(symptoms) == 3
}

func (ph *prescriberHandlerCOVID) OpenPrescription(_ Patient, symptoms []Symptom) (Case, bool) {
	p := NewPrescription(
		"清冠一號",
		"新冠肺炎（專業學名：COVID-19）",
		"清冠一號",
		"將相關藥材裝入茶包裡，使用500 mL 溫、熱水沖泡悶煮1~3 分鐘後即可飲用。",
	)

	return NewCase(
		symptoms,
		p,
		time.Now(),
	), true
}
