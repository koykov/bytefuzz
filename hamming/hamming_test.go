package hamming

import "testing"

var stages = []struct {
	text, target string
	distance     float64
}{
	{"", "", 0},
	{"A", "A", 0},
	{"G", "T", 1},
	{"GGACTGAAATCTG", "GGACTGAAATCTG", 0},
	{"GGACGGATTCTG", "AGGACGGATTCT", 9},
	{"AATG", "AAA", 2},
	{"", "G", 1},
	{"G", "", 1},
}

func TestHamming(t *testing.T) {
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

func BenchmarkHamming(b *testing.B) {
	for _, st := range stages {
		b.Run(st.text, func(b *testing.B) {
			b.ReportAllocs()
			ctx := NewCtx()
			for i := 0; i < b.N; i++ {
				ctx.Reset()
				ctx.DistanceString(st.text, st.target)
			}
		})
	}
}
