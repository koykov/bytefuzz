package jaro_winkler

import "testing"

var stages = []struct {
	text, target string
	distance     float64
}{
	{"apple", "applet", 0.9722222222222223},
	{"Michael", "Michelle", 0.8880952380952382},
	{"McDonald's", "Mcdonells", 0.8777777777777778},
	{"hello", "world", 0.4666666666666666},
	{"telephone", "telephne", 0.9346560846560846},
	{"AI", "Artificial Intelligence", 0},
	{"abacus", "abaxus", 0.8444444444444443},
	{"Thompson", "Thomson", 0.8278571428571428},
}

func TestJaroWinkler(t *testing.T) {
	for _, st := range stages {
		t.Run(st.text, func(t *testing.T) {
			ctx := NewCtx()
			dist := ctx.DistanceString(st.text, st.target)
			if dist != st.distance {
				t.Log(dist)
				t.Errorf("distance: got %f, want %f", dist, st.distance)
			}
		})
	}
}
