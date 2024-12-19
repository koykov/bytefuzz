package cosine_similarity

import (
	"math"

	"github.com/koykov/byteconv"
)

type Ctx struct {
	// todo implement me
}

func NewCtx() *Ctx {
	return &Ctx{}
}

func (ctx *Ctx) Distance(text, target []byte) float64 {
	vect1 := make(map[byte]int)
	for _, t := range text {
		vect1[t]++
	}
	vect2 := make(map[byte]int)
	for _, t := range target {
		vect2[t]++
	}
	// A·B
	dotProduct := 0.0
	for k, v := range vect1 {
		dotProduct += float64(v) * float64(vect2[k])
	}
	// |A|*|B|
	sum1 := 0.0
	for _, v := range vect1 {
		sum1 += math.Pow(float64(v), 2)
	}
	sum2 := 0.0
	for _, v := range vect2 {
		sum2 += math.Pow(float64(v), 2)
	}
	magnitude := math.Sqrt(sum1) * math.Sqrt(sum2)
	if magnitude == 0 {
		return 0.0
	}
	return dotProduct / magnitude
}

func (ctx *Ctx) DistanceString(text, target string) float64 {
	ptext, ptarget := byteconv.S2B(text), byteconv.S2B(target)
	return ctx.Distance(ptext, ptarget)
}

func (ctx *Ctx) Reset() {}
