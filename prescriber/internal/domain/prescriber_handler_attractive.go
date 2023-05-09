package domain

import "time"

type prescriberHandlerAttractive struct{}

func NewPrescriberHandlerAttractive() PrescriberHandler {
	return &prescriberHandlerAttractive{}
}

func (ph *prescriberHandlerAttractive) Match(patient Patient, symptoms []Symptom) bool {
	if patient.Gender == "female" {
		return len(symptoms) == 1 && symptoms[0] == SymptomSneeze
	}

	return false
}

func (ph *prescriberHandlerAttractive) OpenPrescription(patient Patient, symptoms []Symptom) (Case, bool) {
	p := NewPrescription(
		"青春抑制劑",
		"有人想你了 (專業學名：Attractive)",
		"假鬢角、臭味",
		"把假鬢角黏在臉的兩側，讓自己異性緣差一點，自然就不會有人想妳了。",
	)

	return NewCase(
		symptoms,
		p,
		time.Now(),
	), true
}
