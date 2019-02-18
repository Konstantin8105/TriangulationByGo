package tp

import "math"

//TODO нужна структура 1 структура и доипозон высот
type Seacher struct {
	beginTriangle []*TriangleStructure
	pointArray    []*Point
}

func CreateSeachers(triangleStructure *TriangleStructure, box *borderBox, pointArray []*Point) *Seacher {
	maxSize := AMOUNT_SEARCHER_FACTOR
	maxSize *= math.Sqrt(float64(len(pointArray)))
	var amount uint64
	if maxSize > 1.01 {
		amount = uint64(maxSize)
	} else {
		amount = 1
	}

	seacher := &(Seacher{nil, pointArray})
	seacher.beginTriangle = make([]*TriangleStructure, amount)
	var i uint64
	for i = 0; i < amount; i++ {
		seacher.beginTriangle[i] = triangleStructure
	}

	// double heightStep = (box.getY_max() - box.getY_min()) / (double) searcher.length;
	// elevations = new double[searcher.length];
	// for (int i = 0; i < elevations.length; i++) {
	//     elevations[i] = box.getY_min() + i * heightStep;
	// }
	// }
	return seacher
}

// TODO необходимо добавить где-то обработку если не существует
func (seacher *Seacher) movingByConvexHull(point *Point, beginTriangle *TriangleStructure, pointArray []*Point) POINT_TRIANGLE_STATE {
	var value [3]POINT_LINE_STATE
	calc := func(indexPoint1 uint64, indexPoint2 uint64, point *Point) POINT_LINE_STATE {
		return calculateValuePointOnLine(pointArray[indexPoint1], pointArray[indexPoint2], point)
	}
	dist := func(indexPoint1 uint64, indexPoint2 uint64, point *Point) float64 {
		return distanceLineAndPoint(pointArray[indexPoint1], pointArray[indexPoint2], point)
	}
	for {
		value[0] = calc(beginTriangle.IndexNodes[0], beginTriangle.IndexNodes[1], point)
		if isAtRightOfByPOINT_LINE_STATE(value[0]) {
			beginTriangle = beginTriangle.Triangles[0]
		} else {
			var whichOp uint8 = 0
			value[1] = calc(beginTriangle.IndexNodes[1], beginTriangle.IndexNodes[2], point)
			if isAtRightOfByPOINT_LINE_STATE(value[1]) {
				whichOp += 1
			}
			value[2] = calc(beginTriangle.IndexNodes[2], beginTriangle.IndexNodes[0], point)
			if isAtRightOfByPOINT_LINE_STATE(value[2]) {
				whichOp += 2
			}
			if whichOp == 0 {
				break
			} else if whichOp == 1 {
				beginTriangle = beginTriangle.Triangles[1]
			} else if whichOp == 2 {
				beginTriangle = beginTriangle.Triangles[2]
			} else {
				if dist(beginTriangle.IndexNodes[1], beginTriangle.IndexNodes[2], point) >
					dist(beginTriangle.IndexNodes[2], beginTriangle.IndexNodes[0], point) {
					beginTriangle = beginTriangle.Triangles[1]
				} else {
					beginTriangle = beginTriangle.Triangles[2]
				}
			}
		}
	}
	trianglePoints := [3]*Point{
		pointArray[beginTriangle.IndexNodes[0]],
		pointArray[beginTriangle.IndexNodes[1]],
		pointArray[beginTriangle.IndexNodes[2]],
	}
	return statePointInTriangle(point, trianglePoints, value)
}
