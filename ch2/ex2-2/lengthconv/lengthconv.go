package lengthconv

import "fmt"

// Meters
type Meter float64

// Feet
type Feet float64

const (
	// FtToMRatio
	FtToMRatio = 0.3048
	// MtoFtRatio
	MtoFtRatio = 3.28084
)

// FtToM converts feet to meters.
func FtToM(ft Feet) Meter {
	return Meter(ft * FtToMRatio)
}

// MToFt converts meters to feet.
func MToFt(m Meter) Feet {
	return Feet(m * MtoFtRatio)
}

func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
func (ft Feet) String() string { return fmt.Sprintf("%gft", ft) }
