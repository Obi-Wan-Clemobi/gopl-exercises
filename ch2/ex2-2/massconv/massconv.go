package massconv

import "fmt"

// Kilograms
type Kilogram float64

// Pounds
type Pound float64

const (
	KgToLbRatio = 2.20462
	LbToKgRatio = 0.453592
)

// KgsToPounds converts kgs to lbs.
func KgToPound(kgs Kilogram) Pound {
	return Pound(kgs * KgToLbRatio)
}

// LbsToKgs converts lbs to kgs.
func LbsToKgs(lbs Pound) Kilogram {
	return Kilogram(lbs * LbToKgRatio)
}

func (kg Kilogram) String() string { return fmt.Sprintf("%gkg", kg) }
func (lb Pound) String() string    { return fmt.Sprintf("%glb", lb) }
