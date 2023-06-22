package domain

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type PrescribeFacade struct {
	prescriber              *Prescriber
	patientDatabaseFileName string
	prescriptionFileName    string
}

func NewPrescribeFacade(patientDatabaseFileName string, prescriptionFileName string) *PrescribeFacade {
	pf := &PrescribeFacade{
		patientDatabaseFileName: patientDatabaseFileName,
		prescriptionFileName:    prescriptionFileName,
	}

	database := NewPatientDatabase()
	pf.loadPatientDatabase(database)
	pf.prescriber = NewPrescriber(
		database,
		pf.loadSupportedPrescriptionHandler(),
	)

	return pf
}

func (pf *PrescribeFacade) Prescribe(id string, symptoms []Symptom, cb func(pf *PrescribeFacade, c Case)) {
	c, success := pf.prescriber.Prescribe(id, symptoms)
	if !success {
		fmt.Println("找不到症狀！")
		return
	}

	pf.savePatientDatabase()

	cb(pf, c)
}

func (pf *PrescribeFacade) Export(format, filePath string, c Case) {
	if format == "json" {
		jsonFile, err := os.OpenFile(fmt.Sprintf("%s%s", filePath, ".json"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			fmt.Println(err)
		}
		defer func() {
			_ = jsonFile.Close()
		}()
		_ = json.NewEncoder(jsonFile).Encode(c)
	}
	if format == "csv" {
		// TODO: 暫時不實作
	}
}

func (pf *PrescribeFacade) loadSupportedPrescriptionHandler() PrescriberHandler {
	file, err := os.Open(pf.prescriptionFileName)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		_ = file.Close()
	}()

	var handler PrescriberHandler

	r := csv.NewReader(file)
	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		if record[0] == "COVID-19" {
			handler = NewPrescriberHandler(NewPrescriberHandlerCOVID(), handler)
		}
		if record[0] == "Attractive" {
			handler = NewPrescriberHandler(NewPrescriberHandlerAttractive(), handler)
		}
		if record[0] == "SleepApneaSyndrome" {
			handler = NewPrescriberHandler(NewPrescriberHandlerSleepApneaSyndrome(), handler)
		}
	}

	return handler
}

func (pf *PrescribeFacade) loadPatientDatabase(database *PatientDatabase) {
	jsonFile, err := os.Open(pf.patientDatabaseFileName)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		_ = jsonFile.Close()
	}()

	var patients []Patient
	byteValue, _ := io.ReadAll(jsonFile)
	_ = json.Unmarshal(byteValue, &patients)

	for _, patient := range patients {
		database.AddPatient(patient)
	}
}

func (pf *PrescribeFacade) savePatientDatabase() {
	jsonFile, err := os.OpenFile(pf.patientDatabaseFileName, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		_ = jsonFile.Close()
	}()
	patients := pf.prescriber.database.GetAllPatients()
	_ = json.NewEncoder(jsonFile).Encode(patients)
}
