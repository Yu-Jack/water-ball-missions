package domain

type Prescriber struct {
	database *PatientDatabase
	handler  PrescriberHandler
}

func NewPrescriber(database *PatientDatabase, handler PrescriberHandler) *Prescriber {
	return &Prescriber{
		database: database,
		handler:  handler,
	}
}

func (p *Prescriber) Prescribe(id string, symptoms []Symptom) (Case, bool) {
	patient, found := p.database.GetPatient(id)

	if !found {
		return Case{}, false
	}

	c, success := p.handler.OpenPrescription(patient, symptoms)
	if !success {
		return Case{}, false
	}

	p.database.UpdatePatientCase(patient, c)

	return c, true
}

func (p *Prescriber) addHandler(handler PrescriberHandler) {
	p.handler = NewPrescriberHandler(handler, p.handler)
}
