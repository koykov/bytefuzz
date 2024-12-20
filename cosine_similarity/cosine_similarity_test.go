package cosine_similarity

import "testing"

var stages = []struct {
	text, target string
	distance     float64
}{
	{"I love LA and New York", "I love New York and LA", 1.0},
	{"I love LA and New York", "string similarity test...", 0.3108947062827684},
}

func TestCosineSimilarity(t *testing.T) {
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

func BenchmarkCosineSimilarity(b *testing.B) {
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
