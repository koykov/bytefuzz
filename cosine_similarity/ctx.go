package cosine_similarity

import (
	"math"
	"unsafe"

	"github.com/koykov/byteconv"
	"github.com/koykov/openrt"
)

type Ctx struct {
	vec [math.MaxUint8 * 2]uint64
}

func NewCtx() *Ctx {
	return &Ctx{}
}

func (ctx *Ctx) Distance(text, target []byte) float64 {
	for i := 0; i < len(text); i++ {
		ctx.vec[i]++
	}
	for i := 0; i < len(target); i++ {
		ctx.vec[math.MaxUint8+i]++
	}
	vec0, vec1 := ctx.vec[:math.MaxUint8], ctx.vec[math.MaxUint8:]
	// A·B
	var dotProduct float64
	for i := 0; i < math.MaxUint8; i++ {
		dotProduct += float64(vec0[i] * vec1[i])
	}
	// |A|*|B|
	var sum1 float64
	for i := 0; i < math.MaxUint8; i++ {
		sum1 += math.Pow(float64(vec0[i]), 2)
	}
	var sum2 float64
	for i := 0; i < math.MaxUint8; i++ {
		sum2 += math.Pow(float64(vec1[i]), 2)
	}
	magnitude := math.Sqrt(sum1) * math.Sqrt(sum2)
	if magnitude == 0 {
		return 0
	}
	return dotProduct / magnitude
}

func (ctx *Ctx) DistanceString(text, target string) float64 {
	ptext, ptarget := byteconv.S2B(text), byteconv.S2B(target)
	return ctx.Distance(ptext, ptarget)
}

func (ctx *Ctx) Reset() {
	openrt.MemclrUnsafe(unsafe.Pointer(&ctx.vec[0]), math.MaxUint8*16)
}
