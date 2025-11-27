package levenshtein

import (
	"math"
	"unsafe"

	"github.com/koykov/byteconv"
	"github.com/koykov/simd/memclr64"
)

const (
	costIns  = 1
	costDel  = 1
	costRepl = 2
)

type Ctx struct {
	text, target []rune

	mx  [][]int32
	buf []int32
}

func NewCtx() *Ctx {
	return &Ctx{}
}

func (ctx *Ctx) Distance(text, target []byte) float64 {
	ctx.text = byteconv.AppendB2R(ctx.text[:0], text)
	ctx.target = byteconv.AppendB2R(ctx.target[:0], target)
	return ctx.dist(ctx.text, ctx.target)
}

func (ctx *Ctx) DistanceString(text, target string) float64 {
	ctx.text = byteconv.AppendS2R(ctx.text[:0], text)
	ctx.target = byteconv.AppendS2R(ctx.target[:0], target)
	return ctx.dist(ctx.text, ctx.target)
}

func (ctx *Ctx) dist(text, target []rune) float64 {
	w, h := len(target)+1, len(text)+1

	if len(ctx.buf) < 2*w {
		ctx.buf = make([]int32, 2*w)
	}
	for i := 0; i < 2; i++ {
		ctx.mx = append(ctx.mx, ctx.buf[i*w:(i+1)*w])
		ctx.mx[i][0] = int32(i * costDel)
	}
	for i := 1; i < w; i++ {
		ctx.mx[0][i] = int32(i * costIns)
	}

	for i := 1; i < h; i++ {
		c, p := ctx.mx[i%2], ctx.mx[(i-1)%2]
		c[0] = int32(i * costDel)
		for j := 1; j < w; j++ {
			dc := p[j] + costDel
			sc := p[j-1]
			if ctx.text[i-1] != ctx.target[j-1] {
				sc += costRepl
			}
			ic := c[j-1] + costIns
			c[j] = min3(dc, sc, ic)
		}
	}
	return float64(ctx.mx[(h-1)%2][w-1])
}

func (ctx *Ctx) Reset() {
	ctx.mx = ctx.mx[:0]
	if len(ctx.buf) > 0 {
		memclr64.ClearUnsafe(unsafe.Pointer(&ctx.buf[0]), len(ctx.buf)*4)
	}
	ctx.text = ctx.text[:0]
	ctx.target = ctx.target[:0]
}

func min3(a, b, c int32) int32 {
	m := int32(math.MaxInt32)
	if a < m {
		m = a
	}
	if b < m {
		m = b
	}
	if c < m {
		m = c
	}
	return m
}
