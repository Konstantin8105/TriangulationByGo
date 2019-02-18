package tp

import (
	"fmt"

	"github.com/Konstantin8105/tp"
)

func main() {
	fmt.Println("Hello World")
	point := tp.Point{-1.0, -1.0}
	fmt.Println(point)
	coordinates := []tp.Point{
		tp.Point{X: 0.0, Y: 0.0},
		tp.Point{X: 100.0, Y: 100.0},
		tp.Point{X: 200.0, Y: 200.0},
		tp.Point{X: 300.0, Y: 300.0},
		tp.Point{X: 400.0, Y: 400.0},
		tp.Point{X: 500.0, Y: 500.0},
		tp.Point{X: 100.0, Y: 0.0},
	}
	// triangulation := tp.Run(coordinates)
	//
	// fmt.Println("Trinaguation result:")
	// fmt.Printf("%t\n", triangulation)

	fmt.Println("BorderBox result:")
	box := tp.CreateBorderBox()
	for _, p := range coordinates {
		box.AddPoint(*p)
	}
	fmt.Println(box)
	fmt.Printf("Point:\n%t\n", box.insideBox(point))
}
