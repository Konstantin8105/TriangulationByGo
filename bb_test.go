package tp

import (
	"fmt"
	"os"
	"testing"
)

func Test(t *testing.T) {
	b := NewBB()
	b.Add(Point{})
	b.Add(Point{X: 1, Y: 1})
	if !b.Inside(Point{0.5, 0.5}) {
		t.Errorf("is not inside")
	}
}

func Example() {
	b := NewBB()
	b.Add(Point{})
	b.Add(Point{X: 1, Y: 1})
	fmt.Fprintf(os.Stdout, "%s", b)
	// Output:
	// BorderBox
	// x={  0.0000,  1.0000}
	// y={  0.0000,  1.0000}
}
