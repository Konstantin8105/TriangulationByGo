package tp

import "math"

type POINT_LINE_STATE uint8

const (
	RESULT_IS_LESS_ZERO POINT_LINE_STATE = iota
	RESULT_IS_ZERO
	RESULT_IS_MORE_ZERO
)

func calculateDouble(p1 *Point, p2 *Point, p3 *Point) float64 {
	return (p2.Y-p1.Y)*(p3.X-p2.X) - (p3.Y-p2.Y)*(p2.X-p1.X)
}

// if return -1 - result is less 0
// if return 0 - result is 0
// if return 1 - result is more 0
func calculateValuePointOnLine(p1 *Point, p2 *Point, p3 *Point) POINT_LINE_STATE {
	value := calculateDouble(p1, p2, p3)
	if value > precisionEpsilon() {
		return RESULT_IS_MORE_ZERO
	}
	if math.Abs(value) > precisionEpsilon() {
		return RESULT_IS_LESS_ZERO
	}
	return RESULT_IS_ZERO
}

func is3pointsCollinear(p1 *Point, p2 *Point, p3 *Point) bool {
	return calculateValuePointOnLine(p1, p2, p3) == RESULT_IS_ZERO
}

func isCounterClockwise(a *Point, b *Point, c *Point) bool {
	return calculateValuePointOnLine(a, b, c) == RESULT_IS_MORE_ZERO
}

func isAtRightOf(a *Point, b *Point, c *Point) bool {
	return isCounterClockwise(a, b, c)
}

func is3pointsCollinearByPOINT_LINE_STATE(pol POINT_LINE_STATE) bool {
	return pol == RESULT_IS_ZERO
}

func isCounterClockwiseByPOINT_LINE_STATE(pol POINT_LINE_STATE) bool {
	return pol == RESULT_IS_MORE_ZERO
}

func isAtRightOfByPOINT_LINE_STATE(pol POINT_LINE_STATE) bool {
	return isCounterClockwiseByPOINT_LINE_STATE(pol)
}

func distanceLineAndPoint(lineP1 *Point, lineP2 *Point, p *Point) float64 {
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

func isPointInCircle(circlePoints []Point, point *Point) bool {
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
	return result > precisionEpsilon()
}

func isPointInRectangle(point *Point, list ...*Point) bool {
	borderBox := createBorderBox()
	for index, p := range list {
		borderBox.addPoint(*p)
		if index > 2 && borderBox.insideBox(point) {
			return true
		}
	}
	return borderBox.insideBox(point)
}

type POINT_TRIANGLE_STATE uint8

const (
	POINT_LINE_STATE_0 POINT_TRIANGLE_STATE = iota
	POINT_LINE_STATE_1
	POINT_LINE_STATE_2
	POINT_ON_CORNER
	POINT_INSIDE
	POINT_OUTSIDE_LINE_0
	POINT_OUTSIDE_LINE_1
	POINT_OUTSIDE_LINE_2
)

func statePointInTriangle(p *Point,
	trianglePoints [3]*Point,
	values [3]POINT_LINE_STATE) POINT_TRIANGLE_STATE {

	for _, t := range trianglePoints {
		if p.equals(t) {
			return POINT_ON_CORNER
		}
	}

	if isPointInRectangle(p, trianglePoints[0], trianglePoints[1]) {
		if is3pointsCollinearByPOINT_LINE_STATE(values[0]) {
			return POINT_LINE_STATE_0
		}
	}
	if isAtRightOfByPOINT_LINE_STATE(values[0]) {
		return POINT_OUTSIDE_LINE_0
	}

	if isPointInRectangle(p, trianglePoints[1], trianglePoints[2]) {
		if is3pointsCollinearByPOINT_LINE_STATE(values[1]) {
			return POINT_LINE_STATE_1
		}
	}
	if isAtRightOfByPOINT_LINE_STATE(values[1]) {
		return POINT_OUTSIDE_LINE_1
	}

	if isPointInRectangle(p, trianglePoints[2], trianglePoints[0]) {
		if is3pointsCollinearByPOINT_LINE_STATE(values[2]) {
			return POINT_LINE_STATE_2
		}
	}
	if isAtRightOfByPOINT_LINE_STATE(values[2]) {
		return POINT_OUTSIDE_LINE_2
	}

	return POINT_INSIDE
}
