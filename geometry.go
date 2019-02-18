package tp

import (
	"math"

	"github.com/Konstantin8105/tp/point"
)

func eps() float64 {
	return 1e-10
}

type pointLineState uint8

const (
	resultIsLessZero pointLineState = iota
	resultIsZero
	resultIsMoreZero
)

func calculateDouble(p1, p2, p3 point.Point) float64 {
	return (p2.Y-p1.Y)*(p3.X-p2.X) - (p3.Y-p2.Y)*(p2.X-p1.X)
}

func calculateValuepointOnLine(p1, p2, p3 point.Point) pointLineState {
	value := calculateDouble(p1, p2, p3)
	if value > eps() {
		return resultIsMoreZero
	}
	if math.Abs(value) > eps() {
		return resultIsLessZero
	}
	return resultIsZero
}

func distanceLineAndPoint(lineP1 point.Point, lineP2 point.Point, p point.Point) float64 {
	var (
		A        float64
		B        float64 = 1
		C        float64
		distance float64
	)
	if math.Abs(lineP2.Y-lineP1.Y) < math.Abs(lineP2.X-lineP1.X) {
		A = -(lineP2.Y - lineP1.Y) / (lineP2.X - lineP1.X)
		C = -lineP1.Y - A*lineP1.X
		distance = math.Abs((A*p.X + B*p.Y + C) / math.Sqrt(A*A+B*B))
	} else {
		A = -(lineP2.X - lineP1.X) / (lineP2.Y - lineP1.Y)
		C = -lineP1.X - A*lineP1.Y
		distance = math.Abs((A*p.Y + B*p.X + C) / math.Sqrt(A*A+B*B))
	}
	return distance
}

func det(a [3][3]float64) float64 {
	return a[0][0]*a[1][1]*a[2][2] + a[1][0]*a[2][1]*a[0][2] +
		a[0][1]*a[1][2]*a[2][0] - a[0][2]*a[1][1]*a[2][0] -
		a[0][1]*a[1][0]*a[2][2] - a[1][2]*a[2][1]*a[0][0]
}

func isPointInCircle(circlePoints []point.Point, point *point.Point) bool {
	var (
		x1x float64 = circlePoints[0].X - point.X
		y1y float64 = circlePoints[0].Y - point.Y

		x2x float64 = circlePoints[1].X - point.X
		y2y float64 = circlePoints[1].Y - point.Y

		x3x float64 = circlePoints[2].X - point.X
		y3y float64 = circlePoints[2].Y - point.Y
	)

	var result float64 = det([3][3]float64{
		{x1x*x1x + y1y*y1y, x1x, y1y},
		{x2x*x2x + y2y*y2y, x2x, y2y},
		{x3x*x3x + y3y*y3y, x3x, y3y},
	})
	return result > eps()
}

type pointTriangleState uint8

const (
	pointOnLine0 pointTriangleState = iota
	pointOnLine1
	pointOnLine2
	pointOnCorner
	pointInside
	pointOutside
	pointOutsideLine0
	pointOutsideLine1
	pointOutsideLine2
)

func isNear(p1, p2 point.Point) bool {
	return math.Hypot(p1.X-p2.X, p1.Y-p2.Y) < 1e-10
}

func (tr *Triangulation) statePointInTriangle(ip int, tris [3]int) pointTriangleState {
	for i := range tris {
		if isNear(tr.ps[ip], tr.ps[tris[i]]) {
			return pointOnCorner
		}
	}

	ts := []struct {
		trisBegin, trisEnd  int
		onLine, outsideLine pointTriangleState
	}{
		{
			trisBegin:   tris[0],
			trisEnd:     tris[1],
			onLine:      pointOnLine0,
			outsideLine: pointOutsideLine0,
		},
		{
			trisBegin:   tris[1],
			trisEnd:     tris[2],
			onLine:      pointOnLine1,
			outsideLine: pointOutsideLine1,
		},
		{
			trisBegin:   tris[2],
			trisEnd:     tris[0],
			onLine:      pointOnLine2,
			outsideLine: pointOutsideLine2,
		},
	}

	for i := range ts {
		switch calculateValuepointOnLine(
			tr.ps[ip],
			tr.ps[ts[i].trisBegin],
			tr.ps[ts[i].trisEnd],
		) {
		case resultIsZero:
			return ts[i].onLine
		case resultIsMoreZero:
			return ts[i].outsideLine
		}
	}

	return pointInside
}
