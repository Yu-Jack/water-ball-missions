package domain

import (
	"time"
)

type prescriberHandlerSleepApneaSyndrome struct{}

func NewPrescriberHandlerSleepApneaSyndrome() PrescriberHandler {
	return &prescriberHandlerSleepApneaSyndrome{}
}

func (ph *prescriberHandlerSleepApneaSyndrome) Match(patient Patient, symptoms []Symptom) bool {
	height := patient.Height / 100
	if patient.Weight/(height*height) > 26 {
		return len(symptoms) == 1 && symptoms[0] == SymptomSnore
	}

	return false
}

func (ph *prescriberHandlerSleepApneaSyndrome) OpenPrescription(patient Patient, symptoms []Symptom) (Case, bool) {
	p := NewPrescription(
		"打呼抑制劑",
		"睡眠呼吸中止症（專業學名：SleepApneaSyndrome）",
		"一捲膠帶",
		"睡覺時，撕下兩塊膠帶，將兩塊膠帶交錯黏在關閉的嘴巴上，就不會打呼了。",
	)

	return NewCase(
		symptoms,
		p,
		time.Now(),
	), true
}
