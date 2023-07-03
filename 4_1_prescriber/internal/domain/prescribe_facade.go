package domain

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"time"
)

type PrescribeFacade struct {
	prescriber              *Prescriber
	patientDatabaseFileName string
	prescriptionFileName    string
	q                       chan Queue
	wait                    chan struct{}
}

type ClientAction interface {
	Export(format, filePath string, c Case)
}

type Queue struct {
	id       string
	symptoms []Symptom
	cb       func(ca ClientAction, c Case)
}

func NewPrescribeFacade(patientDatabaseFileName string, prescriptionFileName string) *PrescribeFacade {
	pf := &PrescribeFacade{
		patientDatabaseFileName: patientDatabaseFileName,
		prescriptionFileName:    prescriptionFileName,
		q:                       make(chan Queue, 1),
		wait:                    make(chan struct{}),
	}

	database := NewPatientDatabase()
	pf.loadPatientDatabase(database)
	pf.prescriber = NewPrescriber(
		database,
		pf.loadSupportedPrescriptionHandler(),
	)

	return pf
}

func (pf *PrescribeFacade) Open() {
	pf.gracefulShutdown()
	go pf.prescribe()
}

func (pf *PrescribeFacade) Wait() {
	<-pf.wait
}

func (pf *PrescribeFacade) gracefulShutdown() {
	termSignals := make(chan os.Signal)
	signal.Notify(termSignals)

	go func() {
		select {
		case <-termSignals:
			pf.wait <- struct{}{}
			close(pf.q)
		}
	}()
}

func (pf *PrescribeFacade) prescribe() {
	for q := range pf.q {
		t1 := time.Now()
		c, success := pf.prescriber.Prescribe(q.id, q.symptoms)
		if !success {
			fmt.Println("找不到症狀！")
			continue
		}

		pf.savePatientDatabase()

		q.cb(pf, c)
		t2 := time.Now()

		if t2.Sub(t1) < 3*time.Second {
			time.Sleep(3*time.Second - t2.Sub(t1))
		}
	}
}

func (pf *PrescribeFacade) Prescribe(id string, symptoms []Symptom, cb func(ca ClientAction, c Case)) {
	// start queuing
	pf.q <- Queue{
		id:       id,
		symptoms: symptoms,
		cb:       cb,
	}
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
