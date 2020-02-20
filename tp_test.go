package tp

import (
	"fmt"
	"testing"
)

func TestTp(t *testing.T) {
	tcs := []struct {
		ps []Point
	}{
		{
			ps: []Point{
				{X: 0.0, Y: 0.0},
				{X: 1.0, Y: 1.0},
				{X: 1.0, Y: 0.0},
			},
		},
	}

	for i := range tcs {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			tris, err := NewTp(tcs[i].ps...)
			if err != nil {
				t.Error(err)
			}
			t.Logf("%s", *tris)
		})
	}
}
