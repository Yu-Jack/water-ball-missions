//go:generate go-enum -f=$GOFILE --nocase

package domain

// Symptom
/*
ENUM(
Sneeze
Headache
Cough
Snore
)
*/
type Symptom string
