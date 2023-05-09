package domain

type PrescriberHandler interface {
	Match(patient Patient, symptoms []Symptom) bool
	OpenPrescription(patient Patient, symptoms []Symptom) (Case, bool)
}

type prescriberHandler struct {
	handler PrescriberHandler
	next    PrescriberHandler
}

func NewPrescriberHandler(handler, next PrescriberHandler) PrescriberHandler {
	return &prescriberHandler{
		handler: handler,
		next:    next,
	}
}

func (ph *prescriberHandler) Match(patient Patient, symptoms []Symptom) bool {
	// for compiler
	return false
}

func (ph *prescriberHandler) OpenPrescription(patient Patient, symptoms []Symptom) (Case, bool) {
	if ph.handler.Match(patient, symptoms) {
		return ph.handler.OpenPrescription(patient, symptoms)
	} else if ph.next != nil {
		return ph.next.OpenPrescription(patient, symptoms)
	}
	return Case{}, false
}
