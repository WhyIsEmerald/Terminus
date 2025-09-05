package data

import (
	"embed"
	"encoding/json"
	"fmt"
)

//go:embed all:measurements
var content embed.FS
var possibleMeasurements []string = []string{
	"absorbed-dose",
	"acceleration",
	"action",
	"activity-radioactivity",
	"amount-of-substance",
	"angular-momentum",
	"area",
	"capacitance",
	"catalytic-activity",
	"conductance",
	"digital-storage",
	"dynamic-viscosity",
	"electric-charge",
	"electric-current",
	"electric-field-strength",
	"energy",
	"entropy",
	"equivalent-dose",
	"force",
	"frequency",
	"heat-capacity",
	"illuminance",
	"inductance",
	"kinematic-viscosity",
	"length",
	"luminous-flux",
	"luminous-intensity",
	"magnetic-field-strength",
	"magnetic-flux-density",
	"magnetic-flux",
	"mass",
	"momentum",
	"permeability",
	"permittivity",
	"plane-angle",
	"power",
	"pressure",
	"resistance",
	"solid-angle",
	"speed",
	"surface-tension",
	"thermodynamic-temperature",
	"time",
	"torque",
	"voltage",
	"volume",
}

func IsValidMeasurement(measurement string) bool {
	for _, m := range possibleMeasurements {
		if m == measurement {
			return true
		}
	}
	return false
}

func ReadUnits(Measurement string) (map[string]any, error) {
	if !IsValidMeasurement(Measurement) {
		return nil, fmt.Errorf("invalid measurement: %s", Measurement)
	}
	file, err := content.Open(fmt.Sprintf("measurements/%s.json", Measurement))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result map[string]any
	err = json.NewDecoder(file).Decode(&result)
	return result, err
}
