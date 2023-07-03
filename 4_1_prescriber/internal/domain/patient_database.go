package domain

type PatientDatabase struct {
	patients map[string]Patient
}

func NewPatientDatabase() *PatientDatabase {
	return &PatientDatabase{
		patients: make(map[string]Patient),
	}
}

func (pdb *PatientDatabase) AddPatient(patient Patient) {
	pdb.patients[patient.ID] = patient
}

func (pdb *PatientDatabase) GetPatient(id string) (Patient, bool) {
	patient, ok := pdb.patients[id]
	return patient, ok
}

func (pdb *PatientDatabase) GetAllPatients() []Patient {
	var patients []Patient
	for _, patient := range pdb.patients {
		patients = append(patients, patient)
	}
	return patients
}

func (pdb *PatientDatabase) UpdatePatientCase(patient Patient, c Case) {
	patient.Cases = append(patient.Cases, c)
	pdb.patients[patient.ID] = patient
}
