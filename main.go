package main

import (
	"fmt"

	"github.com/Konstantin8105/TriangulationByGo/folder1"
)

func main() {
	fmt.Println("Hello World")
	point := createPointByCoordinate(-1.0, -1.0)
	fmt.Println(point)
	coordinates := []*Point{
		createPointByCoordinate(0.0, 0.0),
		createPointByCoordinate(100.0, 100.0),
		createPointByCoordinate(200.0, 200.0),
		createPointByCoordinate(300.0, 300.0),
		createPointByCoordinate(400.0, 400.0),
		createPointByCoordinate(500.0, 500.0),
		createPointByCoordinate(100.0, 0.0),
	}
	triangulation := createTriangulation(coordinates)

	fmt.Println("Trinaguation result:")
	fmt.Printf("%t\n", triangulation)

	fmt.Println("BorderBox result:")
	box := createBorderBox()
	for _, p := range coordinates {
		box.addPoint(*p)
	}
	fmt.Println(box)
	fmt.Printf("Point:\n%t\n", box.insideBox(point))
	strange.Show()
}
