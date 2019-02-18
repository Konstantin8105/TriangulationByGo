package tp

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func createPointByCoordinate(input_x float64, input_y float64) *Point {
	return &(Point{input_x, input_y})
}

func (p *Point) String() string {
	return fmt.Sprintf("Point(%e,%e)", p.x, p.y)
}

// TODO: use DEEP COMPARE
func (p *Point) equals(point *Point) bool {
	if math.Abs(p.x-point.x) > precisionEpsilon() || math.Abs(p.y-point.y) > precisionEpsilon() {
		return false
	}
	return true
}
