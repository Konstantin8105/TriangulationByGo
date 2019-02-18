package bb

import (
	"fmt"
	"math"
)

// BorderBox is struct of bordex box.
type BorderBox struct {
	x_min, x_max float64
	y_min, y_max float64
	updateCenter bool
	center       tp.Point
}

// New return prepared border box.
func New() *BorderBox {
	p := new(BorderBox)
	p.x_min = +math.MaxFloat64
	p.x_max = -math.MaxFloat64
	p.y_min = +math.MaxFloat64
	p.y_max = -math.MaxFloat64
	return p
}

// Add update border of border box.
func (box *BorderBox) Add(p tp.Point) {
	box.x_max = math.Max(box.x_max, p.X)
	box.x_min = math.Min(box.x_min, p.X)
	box.y_max = math.Max(box.y_max, p.Y)
	box.y_min = math.Min(box.y_min, p.Y)
	box.updateCenter = true
}

// Center return coordinate of center point.
func (box *BorderBox) Center() tp.Point {
	if box.updateCenter {
		box.center = tp.Point{
			X: (box.x_min + box.x_max) / 2.0,
			Y: (box.y_min + box.y_max) / 2.0,
		}
		box.updateCenter = false
		return box.center
	}
	return box.center
}

// Inside return true if point `p` is inside border box, but not on border.
// Or another case return false.
func (box *BorderBox) Inside(p Point) bool {
	if point.X < box.x_min ||
		point.X > box.x_max ||
		point.Y < box.y_min ||
		point.Y > box.y_max {
		return false
	}
	return true
}

// String return string with border box coordinates
func (box *BorderBox) String() string {
	return fmt.Sprintf("BorderBox\nx={%8.4f,%8.4f}\ny={%8.4f,%8.4f}",
		box.x_min,
		box.x_max,
		box.y_min,
		box.y_max)
}
