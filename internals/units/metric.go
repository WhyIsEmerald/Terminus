package units

var prefixes = map[string]float64{
	"Q":  1e30,
	"R":  1e27,
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
	"":   1, // base unit
	"d":  1e-1,
	"c":  1e-2,
	"m":  1e-3,
	"µ":  1e-6,
	"n":  1e-9,
	"p":  1e-12,
	"f":  1e-15,
	"a":  1e-18,
	"z":  1e-21,
	"y":  1e-24,
	"r":  1e-27,
	"q":  1e-30,
}
var baseUnits = map[string]string{
	"length":      "m",
	"mass":        "kg",
	"time":        "s",
	"current":     "A",
	"temperature": "K",
	"amount":      "mol",
	"luminous":    "cd",
	"angle":       "rad",
}

type DerivedUnit struct {
	symbol       string
	value        string
	hasOwnSymbol bool
}

var derivedUnits = map[string]DerivedUnit{
	// Geometric
	"area":   {symbol: "m²", value: "m²", hasOwnSymbol: false},
	"volume": {symbol: "m³", value: "m³", hasOwnSymbol: false},

	// Kinematics
	"speed":           {symbol: "m/s", value: "m/s", hasOwnSymbol: false},
	"acceleration":    {symbol: "m/s²", value: "m/s²", hasOwnSymbol: false},
	"momentum":        {symbol: "kg·m/s", value: "kg*m/s", hasOwnSymbol: false},
	"angularMomentum": {symbol: "J·s", value: "kg*m²/s", hasOwnSymbol: false},
	"action":          {symbol: "J·s", value: "kg*m²/s", hasOwnSymbol: false},

	// Mechanics
	"force":          {symbol: "N", value: "kg*m/s²", hasOwnSymbol: true},
	"pressure":       {symbol: "Pa", value: "kg/(m*s²)", hasOwnSymbol: true},
	"torque":         {symbol: "N·m", value: "kg*m²/s²", hasOwnSymbol: false},
	"surfaceTension": {symbol: "N/m", value: "kg/s²", hasOwnSymbol: false},

	// Energy & Power
	"energy":               {symbol: "J", value: "kg*m²/s²", hasOwnSymbol: true},
	"power":                {symbol: "W", value: "kg*m²/s³", hasOwnSymbol: true},
	"heatCapacity":         {symbol: "J/K", value: "kg*m²/(s²*K)", hasOwnSymbol: false},
	"specificHeatCapacity": {symbol: "J/(kg·K)", value: "m²/(s²*K)", hasOwnSymbol: false},
	"molarHeatCapacity":    {symbol: "J/(mol·K)", value: "kg*m²/(s²*mol*K)", hasOwnSymbol: false},
	"entropy":              {symbol: "J/K", value: "kg*m²/(s²*K)", hasOwnSymbol: false},

	// Electricity & Magnetism
	"charge":        {symbol: "C", value: "A*s", hasOwnSymbol: true},
	"voltage":       {symbol: "V", value: "kg*m²/(s³*A)", hasOwnSymbol: true},
	"capacitance":   {symbol: "F", value: "s⁴*A²/(kg*m²)", hasOwnSymbol: true},
	"resistance":    {symbol: "Ω", value: "kg*m²/(s³*A²)", hasOwnSymbol: true},
	"conductance":   {symbol: "S", value: "s³*A²/(kg*m²)", hasOwnSymbol: true},
	"inductance":    {symbol: "H", value: "kg*m²/(s²*A²)", hasOwnSymbol: true},
	"electricField": {symbol: "V/m", value: "kg*m/(s³*A)", hasOwnSymbol: false},
	"permittivity":  {symbol: "F/m", value: "A²*s⁴/(kg*m³)", hasOwnSymbol: false},
	"permeability":  {symbol: "H/m", value: "kg*m/(s²*A²)", hasOwnSymbol: false},

	// Magnetism
	"magneticFlux":          {symbol: "Wb", value: "kg*m²/(s²*A)", hasOwnSymbol: true},
	"magneticFluxDensity":   {symbol: "T", value: "kg/(s²*A)", hasOwnSymbol: true},
	"magneticFieldStrength": {symbol: "A/m", value: "A/m", hasOwnSymbol: false},

	// Waves & Radiation
	"frequency":      {symbol: "Hz", value: "1/s", hasOwnSymbol: true},
	"activity":       {symbol: "Bq", value: "1/s", hasOwnSymbol: true},
	"absorbedDose":   {symbol: "Gy", value: "m²/s²", hasOwnSymbol: true},
	"equivalentDose": {symbol: "Sv", value: "m²/s²", hasOwnSymbol: true},

	// Luminous
	"luminousFlux": {symbol: "lm", value: "cd", hasOwnSymbol: true},
	"illuminance":  {symbol: "lx", value: "cd/m²", hasOwnSymbol: true},

	// Thermodynamics
	"catalyticActivity": {symbol: "kat", value: "mol/s", hasOwnSymbol: true},

	// Dimensionless
	"angle":      {symbol: "rad", value: "1", hasOwnSymbol: true},
	"solidAngle": {symbol: "sr", value: "1", hasOwnSymbol: true},

	// Fluid properties
	"dynamicViscosity":   {symbol: "Pa·s", value: "kg/(m*s)", hasOwnSymbol: false},
	"kinematicViscosity": {symbol: "m²/s", value: "m²/s", hasOwnSymbol: false},
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
