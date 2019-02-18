package main

import (
	"fmt"

	"github.com/Konstantin8105/tp/bb"
	"github.com/Konstantin8105/tp/point"
)

func main() {
	p := point.Point{-1.0, -1.0}
	fmt.Println(p)
	coordinates := []point.Point{
		point.Point{X: 0.0, Y: 0.0},
		point.Point{X: 100.0, Y: 100.0},
		point.Point{X: 200.0, Y: 200.0},
		point.Point{X: 300.0, Y: 300.0},
		point.Point{X: 400.0, Y: 400.0},
		point.Point{X: 500.0, Y: 500.0},
		point.Point{X: 100.0, Y: 0.0},
	}
	// triangulation := tp.Run(coordinates)
	//
	// fmt.Println("Trinaguation result:")
	// fmt.Printf("%t\n", triangulation)

	fmt.Println("BorderBox result:")
	box := bb.New()
	for _, p := range coordinates {
		box.Add(p)
	}
	fmt.Println(box)
	fmt.Printf("Point:\n%t\n", box.Inside(p))
}
