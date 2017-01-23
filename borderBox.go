package main

import (
	"fmt"
	"math"
)

type BorderBox struct {
	x_min, x_max    float64
	y_min, y_max    float64
	calculateCenter bool
	center          *Point
}

func createBorderBox() *BorderBox {
	p := new(BorderBox)
	p.x_min = math.MaxFloat64
	p.x_max = -math.MaxFloat64
	p.y_min = math.MaxFloat64
	p.y_max = -math.MaxFloat64
	return p
}

func (box *BorderBox) addPoint(point Point) {
	box.x_max = math.Max(box.x_max, point.x)
	box.x_min = math.Min(box.x_min, point.x)
	box.y_max = math.Max(box.y_max, point.y)
	box.y_min = math.Min(box.y_min, point.y)
	box.calculateCenter = true
}

func (box *BorderBox) getCenter() *Point {
	if box.calculateCenter {
		box.center = createPointByCoordinate((box.x_min+box.x_max)/2.0, (box.y_min+box.y_max)/2.0)
		box.calculateCenter = false
		return box.center
	}
	return box.center
}

func (box *BorderBox) insideBox(point *Point) bool {
	if point.x < box.x_min || point.x > box.x_max || point.y < box.y_min || point.y > box.y_max {
		return false
	}
	return true
}

func (box *BorderBox) String() string {
	return fmt.Sprintf("BorderBox\nx={%e,%e}\ny={%e,%e}", box.x_min, box.x_max, box.y_min, box.y_max)
}
