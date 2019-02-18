package tp

import (
	"fmt"
	"math"
)

type borderBox struct {
	x_min, x_max    float64
	y_min, y_max    float64
	calculateCenter bool
	center          *Point
}

func createBorderBox() *borderBox {
	p := new(borderBox)
	p.x_min = +math.MaxFloat64
	p.x_max = -math.MaxFloat64
	p.y_min = +math.MaxFloat64
	p.y_max = -math.MaxFloat64
	return p
}

func (box *borderBox) addPoint(point Point) {
	box.x_max = math.Max(box.x_max, point.X)
	box.x_min = math.Min(box.x_min, point.X)
	box.y_max = math.Max(box.y_max, point.Y)
	box.y_min = math.Min(box.y_min, point.Y)
	box.calculateCenter = true
}

func (box *borderBox) getCenter() *Point {
	if box.calculateCenter {
		box.center = &Point{
			X: (box.x_min + box.x_max) / 2.0,
			Y: (box.y_min + box.y_max) / 2.0,
		}
		box.calculateCenter = false
		return box.center
	}
	return box.center
}

func (box *borderBox) insideBox(point *Point) bool {
	if point.X < box.x_min || point.X > box.x_max || point.Y < box.y_min || point.Y > box.y_max {
		return false
	}
	return true
}

func (box *borderBox) String() string {
	return fmt.Sprintf("BorderBox\nx={%e,%e}\ny={%e,%e}", box.x_min, box.x_max, box.y_min, box.y_max)
}
