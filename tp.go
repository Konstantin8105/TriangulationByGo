package tp

import (
	"container/list"
	"fmt"

	"github.com/Konstantin8105/tp/bb"
	"github.com/Konstantin8105/tp/point"
)

// main data structure for triangulation
//
// Triangulation data structure  "Nodes, simple ribs и triangles"
// book "Algoritm building and analyse triangulation", A.B.Skvorcov.
// paragraph 1.2.5
//
type data struct {
	nodes     [3]int   // indexes of triangle points
	triangles [3]*data // indexes of near triangles
	ribs      [3]int   // indexes of triangle ribs
}

func (d *data) changeClockwise() {
	d.nodes[0], d.nodes[1] = d.nodes[1], d.nodes[0]
	d.ribs[1], d.ribs[2] = d.ribs[2], d.ribs[1]
	d.triangles[1], d.triangles[2] = d.triangles[2], d.triangles[1]
}

type Triangulation struct {
	ps []point.Point
	ds *list.List
}

func New(ps ...point.Point) (tr *Triangulation, err error) {
	if len(ps) < 3 {
		err = fmt.Errorf("not enougt input points")
		return
	}
	tr = &Triangulation{}
	tr.ds = list.New()

	// find border box
	b := bb.New()
	for i := range ps {
		b.Add(ps[i])
	}

	//
	// create pseudo-box.
	// all points must be inside pseudo-box
	//
	//	P1     P2
	//	o---1---o
	//	|      /|
	//	|  0  / |
	//	|    /  |
	//	0   2   3
	//	|  /    |
	//	| /  1  |
	//  |/      |
	//  o---4---o
	//	P0     P3
	//
	scale := 20.0
	xSize := b.Xmax - b.Xmin
	ySize := b.Ymax - b.Ymin
	pps := []point.Point{ // pseudo-box points
		point.Point{X: b.Xmin - scale*xSize, Y: b.Ymin - scale*ySize}, // P0
		point.Point{X: b.Xmin - scale*xSize, Y: b.Ymax + scale*ySize}, // P1
		point.Point{X: b.Xmax + scale*xSize, Y: b.Ymax + scale*ySize}, // P2
		point.Point{X: b.Xmax + scale*xSize, Y: b.Ymin - scale*ySize}, // P3
	}
	defer func() {
		for i := range pps {
			tr.remove(pps[i])
		}
	}()
	tr.ps = append(pps, ps...)

	//
	// create points, ribs, triangles pseudo-box
	//
	t0 := &data{
		nodes: [3]int{0, 2, 1},
		ribs:  [3]int{0, 2, 1},
	}
	t1 := &data{
		nodes: [3]int{2, 0, 3},
		ribs:  [3]int{3, 2, 4},
	}
	t0.triangles[2] = t1
	t1.triangles[2] = t0
	tr.ds.PushFront(t0)
	tr.ds.PushFront(t1)

	//
	// add points in triangles
	//
	for i := 5; i < len(tr.ps); i++ {
		if err := tr.add(i); err != nil {
			return tr, err
		}
	}
	return tr, nil
}

func (tr *Triangulation) remove(p point.Point) error {
	return fmt.Errorf("add implementation for remove")
}

func (tr *Triangulation) add(next int) (err error) {
	if debugFlag {
		logger.Printf("add point #%d : %s", next, tr.ps[next])
	}
	state, tri := tr.findTriangle(next)
	err = fmt.Errorf("Strange point #%d with state `%s` : %s", next, state, tr.ps[next].String())
	switch state {
	case pointInside:
		err = tr.addInTriangle(tri, next)
	case pointOnLine0, pointOnLine1, pointOnLine2:
		err = tr.addOnLine(tri, next, state)
	case pointOnCorner:
		err = nil
	}
	return
}

func (tr *Triangulation) findTriangle(next int) (state pointTriangleState, tri *data) {
	var found bool
	for e := tr.ds.Front(); e != nil; e = e.Next() {
		// moving triangle by triangles
		tri = e.Value.(*data)
		state = tr.statePointInTriangle(next, tri.nodes)
		switch state {
		case pointOutsideLine0, pointOutsideLine1, pointOutsideLine2:
			found = false
		default:
			found = true
		}
		if found {
			break
		}
	}
	if !found {
		state = pointOutside
	}
	return
}

func (tr *Triangulation) addOnLine(tri *data, next int, state pointTriangleState) (err error) {
	return fmt.Errorf("add implementation for addOnLine")
}

func (tr *Triangulation) addInTriangle(tri *data, next int) (err error) {
	return fmt.Errorf("add implementation for addInTriangle")
}

func (tr Triangulation) String() string {
	var out string
	out += "some"
	return out
}
