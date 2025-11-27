package cosine_similarity

import (
	"math"
	"unsafe"

	"github.com/koykov/byteconv"
	"github.com/koykov/simd/memclr64"
)

type Ctx struct {
	vec [math.MaxUint8 * 2]float64
	pc  [math.MaxUint8]float64
}

func NewCtx() *Ctx {
	return &Ctx{}
}

func (ctx *Ctx) Distance(text, target []byte) float64 {
	_ = ctx.vec[math.MaxUint8*2-1]
	for i := 0; i < len(text); i++ {
		ctx.vec[text[i]]++
	}
	for i := 0; i < len(target); i++ {
		ctx.vec[math.MaxUint8+int(target[i])]++
	}
	vec0, vec1 := ctx.vec[:math.MaxUint8], ctx.vec[math.MaxUint8:]
	// A·B
	var dotp float64
	_, _ = vec0[math.MaxUint8-1], vec1[math.MaxUint8-1]
	for i := 0; i < math.MaxUint8; i++ {
		dotp += vec0[i] * vec1[i]
	}
	// |A|*|B|
	var sum0 float64
	for i := 0; i < math.MaxUint8; i++ {
		if vec0[i] == 0 {
			continue
		}
		sum0 += ctx.pow2(vec0[i])
	}
	var sum1 float64
	for i := 0; i < math.MaxUint8; i++ {
		if vec1[i] == 0 {
			continue
		}
		sum1 += ctx.pow2(vec1[i])
	}
	mag := math.Sqrt(sum0) * math.Sqrt(sum1)
	if mag == 0 {
		return 0
	}
	return dotp / mag
}

func (ctx *Ctx) DistanceString(text, target string) float64 {
	ptext, ptarget := byteconv.S2B(text), byteconv.S2B(target)
	return ctx.Distance(ptext, ptarget)
}

func (ctx *Ctx) pow2(x float64) float64 {
	_ = ctx.pc[math.MaxUint8-1]
	i := uint8(x)
	if c := ctx.pc[i]; c > 0 {
		return c
	}
	ctx.pc[i] = math.Pow(x, 2)
	return ctx.pc[i]
}

func (ctx *Ctx) Reset() {
	memclr64.ClearUnsafe(unsafe.Pointer(&ctx.vec[0]), math.MaxUint8*16)
}
