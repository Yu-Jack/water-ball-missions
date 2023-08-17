package domain

import (
	"os"
	"os/signal"
	"time"
)

type Prescriber struct {
	database *PatientDatabase
	handler  PrescriberHandler

	q    chan Queue
	wait chan struct{}
}

type Queue struct {
	id       string
	symptoms []Symptom
	result   chan result
}

type result struct {
	Success bool
	Case    Case
}

func NewPrescriber(database *PatientDatabase, handler PrescriberHandler) *Prescriber {
	p := &Prescriber{
		database: database,
		handler:  handler,
		q:        make(chan Queue, 1),
		wait:     make(chan struct{}, 1),
	}

	p.Open()

	return p
}

func (p *Prescriber) Open() {
	p.gracefulShutdown()
	go p.prescribe()
}

// Wait should be outside in normal case
func (p *Prescriber) Wait() {
	<-p.wait
}

func (p *Prescriber) gracefulShutdown() {
	termSignals := make(chan os.Signal)
	signal.Notify(termSignals)

	go func() {
		select {
		case <-termSignals:
			p.wait <- struct{}{}
			close(p.q)
		}
	}()
}

func (p *Prescriber) prescribe() {
	for q := range p.q {
		t1 := time.Now()

		patient, found := p.database.GetPatient(q.id)

		if !found {
			q.result <- result{Success: false}
			continue
		}

		c, success := p.handler.OpenPrescription(patient, q.symptoms)
		if !success {
			q.result <- result{Success: false}
			continue
		}

		p.database.UpdatePatientCase(patient, c)

		q.result <- result{Success: true, Case: c}

		t2 := time.Now()

		if t2.Sub(t1) < 3*time.Second {
			time.Sleep(3*time.Second - t2.Sub(t1))
		}
	}
}

func (p *Prescriber) Prescribe(id string, symptoms []Symptom) (Case, bool) {
	q := Queue{
		id:       id,
		symptoms: symptoms,
		result:   make(chan result),
	}

	p.q <- q

	r := <-q.result

	return r.Case, r.Success
}

func (p *Prescriber) addHandler(handler PrescriberHandler) {
	p.handler = NewPrescriberHandler(handler, p.handler)
}
