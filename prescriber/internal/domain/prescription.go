package domain

type Prescription struct {
	Name             string `json:"name"`
	PotentialDisease string `json:"potential_disease"`
	Medicines        string `json:"medicines"`
	Usage            string `json:"usage"`
}

func NewPrescription(name string, potentialDisease string, medicines string, usage string) Prescription {
	return Prescription{
		Name:             name,
		PotentialDisease: potentialDisease,
		Medicines:        medicines,
		Usage:            usage,
	}
}
