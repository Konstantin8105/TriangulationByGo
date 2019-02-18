package tp

import "container/list"

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
}

func (tr *Triangulation) New(ps ...point.Point) {
	tr.ps = ps

	// find border box
	b := bb.New()
	for i := range ps {
		b.Add(ps[i])
	}

	// create pseudo-box.
	// all points must be inside pseudo-box
	//
	//	P1     P2
	//	o-------o
	//	|      /|
	//	|     / |
	//	|    /  |
	//	|   /   |
	//	|  /    |
	//	| /     |
	//  |/      |
	//  o-------o
	//	P0     P3
	//
	pps := []point.Point{ // pseudo-box points
		point.Point{X: bb.Xmin - 1.0, Y: bb.Ymin - 1.0}, // P0
		point.Point{X: bb.Xmin - 1.0, Y: bb.Ymax + 1.0}, // P1
		point.Point{X: bb.Xmax + 1.0, Y: bb.Ymax + 1.0}, // P2
		point.Point{X: bb.Xmax + 1.0, Y: bb.Ymin - 1.0}, // P3
	}
	defer func() {
		for i := range pps {
			tr.remove(pps[i])
		}
	}()

	// flipper = new FliperDelaunay(this);
	// List<Point>[] pointArray = convexHullDouble(input);
	// if (pointArray == null)
	// return;
	// List<Point> convexPoints = pointArray[0];
	// BorderBox box = createConvexHullTriangles(convexPoints);
	// searcher = new FastSearcher(this, triangleList.getFirstNotNullableElement(), box, pointArray[1].size());
	//
	// if (pointArray[1].size() >= MINIMAL_POINTS_FOR_CLEANING) {
	// int amount = (int) (AMOUNT_CLEANING_FACTOR_TRIANGLE_STRUCTURE * pointArray[1].size());
	// amount = amount < 1 ? 1 : amount;
	// triangleList.setMaxAmountNullableElements(amount);
	// }
	// for (int i = 0; i < pointArray[1].size(); i++) {
	// addNextPoint(pointArray[1].get(i));
	// flipper.run();
	//            if (i % 1000 == 0)
	//                System.err.println(i);
	// }
	// flipper.run();

}

/*
    //Performance O(n*log(n)) in worst case
    // Point[0][] - convex points
    // Point[1][] - sorted list of all points
    private static List[] convexHullDouble(Point[] inputPoints) {
        if (inputPoints.length < 2) {
            return null;
        }

        ArrayList<Point> array = new ArrayList<>(Arrays.asList(inputPoints));

        Collections.sort(array, new Comparator<Point>() {
            @Override
            public int compare(Point first, Point second) {
                if ((first).getX() == (second).getX()) {
                    if ((first).getY() > (second).getY())
                        return 1;
                    if ((first).getY() < (second).getY())
                        return -1;
                    return 0;
                }
                if ((first).getX() > (second).getX())
                    return 1;
                if ((first).getX() < (second).getX())
                    return -1;
                return 0;
            }
        });

        List<Integer> removedIndex = new ArrayList<>();
        for (int i = 1; i < array.size(); i++) {
            if (array.get(i - 1).equals(array.get(i))) {
                removedIndex.add(i);
            }
        }
        for (int i = removedIndex.size() - 1; i >= 0; i--) {
            int position = removedIndex.get(i);
            array.remove(position);
        }

        int n = array.size();
        Point[] P = new Point[n];
        for (int i = 0; i < n; i++) {
            P[i] = array.get(i);
        }

        Point[] H = new Point[2 * n];
//            List<Point> H = new ArrayList<>(2 * n);

        int k = 0;
//             Build lower hull
        for (int i = 0; i < n; ++i) {
            while (k >= 2 && Geometry.isCounterClockwise(H[k - 2], H[k - 1], P[i])) {
                k--;
            }
            H[k++] = P[i];
        }

        // Build upper hull
        for (int i = n - 2, t = k + 1; i >= 0; i--) {
            while (k >= t && Geometry.isCounterClockwise(H[k - 2], H[k - 1], P[i])) {
                k--;
            }
            H[k++] = P[i];
        }
        List<Point> convexPoints = new ArrayList<>();
        if (k > 1) {
            H = Arrays.copyOfRange(H, 0, k - 1); // remove non-hull vertices after k; remove k - 1 which is a duplicate
            boolean[] removed = new boolean[k - 1];
            for (int position0 = 0; position0 < removed.length; position0++) {
                int position1 = position0 + 1 >= removed.length ? position0 + 1 - removed.length : position0 + 1;
                int position2 = position0 + 2 >= removed.length ? position0 + 2 - removed.length : position0 + 2;
                if (Geometry.is3pointsCollinear(
                        H[position0],
                        H[position1],
                        H[position2])) {
                    removed[position1] = true;
                }
            }
            for (int i = 0; i < removed.length; i++) {
                if (!removed[i])
                    convexPoints.add(H[i]);
            }
            if (array.size() > 5) {
                if ((double) convexPoints.size() / (double) array.size() > RATIO_DELETING_CONVEX_POINT_FROM_POINT_LIST) {
                    boolean[] delete = new boolean[array.size()];
                    int position = 0;
                    for (int i = 0; i < array.size(); i++) {
                        if (array.get(i).equals(convexPoints.get(position))) {
                            delete[i] = true;
                            position++;
                        }
                    }
                    for (int i = array.size() - 1; i >= 0; i--) {
                        if (position >= convexPoints.size())
                            break;
                        if (array.get(i).equals(convexPoints.get(position))) {
                            delete[i] = true;
                            position++;
                        }
                    }
                    ArrayList<Point> newList = new ArrayList<>();
                    for (int i = 0; i < array.size(); i++) {
                        if (!delete[i])
                            newList.add(array.get(i));
                    }
                    array = newList;
                }
            }
        }

        return new List[]{convexPoints, array};
    }


*/

type grid = *list.List

func (g *grid) add(ds ...data) {
	if g == nil {
		g = list.New()
	}
	for i := range ds {
		g.PushFront(ds[i])
	}
}
