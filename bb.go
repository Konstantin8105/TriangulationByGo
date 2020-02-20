package tp

import (
	"fmt"
	"math"
)

// BorderBox is struct of bordex box.
type BorderBox struct {
	Xmin, Xmax   float64
	Ymin, Ymax   float64
	updateCenter bool
	center       Point
}

// New return prepared border box.
func NewBB() *BorderBox {
	p := new(BorderBox)
	p.Xmin = +math.MaxFloat64
	p.Xmax = -math.MaxFloat64
	p.Ymin = +math.MaxFloat64
	p.Ymax = -math.MaxFloat64
	return p
}

// Add update border of border box.
func (box *BorderBox) Add(p Point) {
	box.Xmax = math.Max(box.Xmax, p.X)
	box.Xmin = math.Min(box.Xmin, p.X)
	box.Ymax = math.Max(box.Ymax, p.Y)
	box.Ymin = math.Min(box.Ymin, p.Y)
	box.updateCenter = true
}

// Center return coordinate of center
func (box *BorderBox) Center() Point {
	if box.updateCenter {
		box.center = Point{
			X: (box.Xmin + box.Xmax) / 2.0,
			Y: (box.Ymin + box.Ymax) / 2.0,
		}
		box.updateCenter = false
		return box.center
	}
	return box.center
}

// Inside return true if point `p` is inside border box, but not on border.
// Or another case return false.
func (box *BorderBox) Inside(p Point) bool {
	if p.X < box.Xmin ||
		p.X > box.Xmax ||
		p.Y < box.Ymin ||
		p.Y > box.Ymax {
		return false
	}
	return true
}

// String return string with border box coordinates
func (box *BorderBox) String() string {
	return fmt.Sprintf("BorderBox\nx={%8.4f,%8.4f}\ny={%8.4f,%8.4f}",
		box.Xmin,
		box.Xmax,
		box.Ymin,
		box.Ymax)
}
