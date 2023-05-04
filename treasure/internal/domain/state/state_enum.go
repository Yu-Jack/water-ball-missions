// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package state

import (
	"errors"
	"fmt"
	"strings"
)

const (
	// TypeAccelerated is a Type of type Accelerated.
	TypeAccelerated Type = "Accelerated"
	// TypeErupting is a Type of type Erupting.
	TypeErupting Type = "Erupting"
	// TypeHealing is a Type of type Healing.
	TypeHealing Type = "Healing"
	// TypeInvincible is a Type of type Invincible.
	TypeInvincible Type = "Invincible"
	// TypeNormal is a Type of type Normal.
	TypeNormal Type = "Normal"
	// TypeOrderless is a Type of type Orderless.
	TypeOrderless Type = "Orderless"
	// TypePoisoned is a Type of type Poisoned.
	TypePoisoned Type = "Poisoned"
	// TypeStockpile is a Type of type Stockpile.
	TypeStockpile Type = "Stockpile"
	// TypeTeleport is a Type of type Teleport.
	TypeTeleport Type = "Teleport"
)

var ErrInvalidType = errors.New("not a valid Type")

// String implements the Stringer interface.
func (x Type) String() string {
	return string(x)
}

// String implements the Stringer interface.
func (x Type) IsValid() bool {
	_, err := ParseType(string(x))
	return err == nil
}

var _TypeValue = map[string]Type{
	"Accelerated": TypeAccelerated,
	"accelerated": TypeAccelerated,
	"Erupting":    TypeErupting,
	"erupting":    TypeErupting,
	"Healing":     TypeHealing,
	"healing":     TypeHealing,
	"Invincible":  TypeInvincible,
	"invincible":  TypeInvincible,
	"Normal":      TypeNormal,
	"normal":      TypeNormal,
	"Orderless":   TypeOrderless,
	"orderless":   TypeOrderless,
	"Poisoned":    TypePoisoned,
	"poisoned":    TypePoisoned,
	"Stockpile":   TypeStockpile,
	"stockpile":   TypeStockpile,
	"Teleport":    TypeTeleport,
	"teleport":    TypeTeleport,
}

// ParseType attempts to convert a string to a Type.
func ParseType(name string) (Type, error) {
	if x, ok := _TypeValue[name]; ok {
		return x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := _TypeValue[strings.ToLower(name)]; ok {
		return x, nil
	}
	return Type(""), fmt.Errorf("%s is %w", name, ErrInvalidType)
}
