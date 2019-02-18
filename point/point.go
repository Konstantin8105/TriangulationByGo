package point

import (
	"fmt"
	"math"
)

// Point store coordinates of point
type Point struct {
	X, Y float64
}

// String return string with coordinates
func (p *Point) String() string {
	return fmt.Sprintf("Point(%8.4f,%8.4f)", p.X, p.Y)
}

// TODO: use DEEP COMPARE
func (p Point) Equals(point Point) bool {
	if math.Abs(p.X-point.X) > eps() || math.Abs(p.Y-point.Y) > eps() {
		return false
	}
	return true
}

func eps() float64 {
	return 1e-10
}
