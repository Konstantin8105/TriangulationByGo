package tp

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p *Point) String() string {
	return fmt.Sprintf("Point(%e,%e)", p.X, p.Y)
}

// TODO: use DEEP COMPARE
func (p *Point) equals(point *Point) bool {
	if math.Abs(p.X-point.X) > eps() || math.Abs(p.Y-point.Y) > eps() {
		return false
	}
	return true
}
