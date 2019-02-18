package tp_test

import (
	"fmt"
	"testing"

	"github.com/Konstantin8105/tp"
	"github.com/Konstantin8105/tp/point"
)

func Test(t *testing.T) {
	tcs := []struct {
		ps []point.Point
	}{
		{
			ps: []point.Point{
				{X: 0.0, Y: 0.0},
				{X: 1.0, Y: 1.0},
				{X: 1.0, Y: 0.0},
			},
		},
	}

	for i := range tcs {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			tris, err := tp.New(tcs[i].ps...)
			if err != nil {
				t.Error(err)
			}
			t.Logf("%s", *tris)
		})
	}
}
