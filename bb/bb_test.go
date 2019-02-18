package bb

import (
	"fmt"
	"os"
	"testing"

	"github.com/Konstantin8105/tp/point"
)

func Test(t *testing.T) {
	b := New()
	b.Add(point.Point{})
	b.Add(point.Point{X: 1, Y: 1})
	if !b.Inside(point.Point{0.5, 0.5}) {
		t.Errorf("is not inside")
	}
}

func Example() {
	b := New()
	b.Add(point.Point{})
	b.Add(point.Point{X: 1, Y: 1})
	fmt.Fprintf(os.Stdout, "%s", b)
	// Output:
	// BorderBox
	// x={  0.0000,  1.0000}
	// y={  0.0000,  1.0000}
}
