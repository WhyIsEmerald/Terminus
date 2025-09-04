package units

import (
	"fmt"

	"github.com/WhyIsEmerald/Terminus/data"
)

func Convert(Measurement string, initial string, target string, magnitude float64) {
	unitsMap, err := data.ReadUnits(Measurement)
	if err != nil {
		fmt.Println(err)
		return
	}
	if _, ok := unitsMap[initial]; !ok {
		fmt.Printf("Invalid initial unit: %s\n", initial)
		return
	}
	if _, ok := unitsMap[target]; !ok {
		fmt.Printf("Invalid target unit: %s\n", target)
		return
	}
	initialUnit := unitsMap[initial].(float64)
	targetUnit := unitsMap[target].(float64)
	result := magnitude * initialUnit / targetUnit
	fmt.Printf("%f %s = %f %s\n", magnitude, initial, result, target)
}
