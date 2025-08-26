package units

import "fmt"

var distMap = generateUnits("m")

func PrintDistance() {
	fmt.Println(distMap)
}
