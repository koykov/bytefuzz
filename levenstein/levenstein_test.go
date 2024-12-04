package levenstein

import "testing"

type stage struct {
	text, target string
	distance     int
}

var stages = []stage{
	{
		text:     "",
		target:   "a",
		distance: 1,
	},
	{
		text:     "a",
		target:   "aa",
		distance: 1,
	},
	{
		text:     "a",
		target:   "aaa",
		distance: 2,
	},
	{
		text:     "",
		target:   "",
		distance: 0,
	},
	{
		text:     "a",
		target:   "b",
		distance: 2,
	},
	{
		text:     "aaa",
		target:   "aba",
		distance: 2,
	},
	{
		text:     "aaa",
		target:   "ab",
		distance: 3,
	},
	{
		text:     "a",
		target:   "a",
		distance: 0,
	},
	{
		text:     "ab",
		target:   "ab",
		distance: 0,
	},
	{
		text:     "a",
		target:   "",
		distance: 1,
	},
	{
		text:     "aa",
		target:   "a",
		distance: 1,
	},
	{
		text:     "aaa",
		target:   "a",
		distance: 2,
	},
	{
		text:     "kitten",
		target:   "sitting",
		distance: 5,
	},
	{
		text:     "kitten",
		target:   "sitting",
		distance: 3,
	},
	{
		text:     "Orange",
		target:   "Apple",
		distance: 5,
	},
	{
		text:     "ab",
		target:   "bc",
		distance: 2,
	},
	{
		text:     "abd",
		target:   "bec",
		distance: 3,
	},
	{
		text:     "me",
		target:   "meme",
		distance: 4,
	},
}

func TestLevenstein(t *testing.T) {
	for _, st := range stages {
		t.Run(st.text, func(t *testing.T) {
			ctx := NewCtx(1, 1, 2, func(a, b rune) bool { return a == b })
			dist := ctx.DistanceString(st.text, st.target)
			if dist != st.distance {
				t.Errorf("distance: got %d, want %d", dist, st.distance)
			}
		})
	}
}
