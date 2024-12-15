package jaro_winkler

import "testing"

var stages = []struct {
	text, target string
	distance     float64
}{
	{"apple", "applet", 0.9722222222222223},
}

func TestJaroWinkler(t *testing.T) {
	for _, st := range stages {
		t.Run(st.text, func(t *testing.T) {
			ctx := NewCtx()
			dist := ctx.DistanceString(st.text, st.target)
			if dist != st.distance {
				t.Errorf("distance: got %f, want %f", dist, st.distance)
			}
		})
	}
}
