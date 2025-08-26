package units

var prefixes = map[string]float64{
	"Y":  1e24,
	"Z":  1e21,
	"E":  1e18,
	"P":  1e15,
	"T":  1e12,
	"G":  1e9,
	"M":  1e6,
	"k":  1e3,
	"h":  1e2,
	"da": 1e1,
	"":   1,
	"d":  1e-1,
	"c":  1e-2,
	"m":  1e-3,
	"μ":  1e-6,
	"n":  1e-9,
	"p":  1e-12,
	"f":  1e-15,
	"a":  1e-18,
	"z":  1e-21,
	"y":  1e-24,
}

func generateUnits(siUnit string) map[string]float64 {
	units := make(map[string]float64)
	for prefix, factor := range prefixes {
		units[prefix+siUnit] = factor
	}
	return units
}

func Convert(value float64, fromUnit, toUnit string) float64 {
	fromFactor := prefixes[fromUnit[:1]]
	toFactor := prefixes[toUnit[:1]]
	return value * fromFactor / toFactor
}
