package hamming

import (
	"github.com/koykov/byteconv"
	"github.com/koykov/bytefuzz"
)

type ctx struct {
	buf []rune
}

func NewCtx() bytefuzz.Interface {
	return &ctx{}
}

func (ctx *ctx) Distance(text, target []byte) float64 {
	ctx.buf = byteconv.AppendB2R(ctx.buf, text)
	r1 := ctx.buf
	ctx.buf = byteconv.AppendB2R(r1, target)
	r2 := ctx.buf[len(r1):]
	return ctx.dist(r1, r2)
}

func (ctx *ctx) DistanceString(text, target string) float64 {
	ctx.buf = byteconv.AppendS2R(ctx.buf, text)
	r1 := ctx.buf
	ctx.buf = byteconv.AppendS2R(r1, target)
	r2 := ctx.buf[len(r1):]
	return ctx.dist(r1, r2)
}

func (ctx *ctx) dist(r1, r2 []rune) (d float64) {
	if len(r2) > len(r1) {
		r1, r2 = r2, r1
	}
	d = float64(len(r1) - len(r2))
	for i := 0; i < len(r2); i++ {
		if r1[i] != r2[i] {
			d++
		}
	}
	return
}

func (ctx *ctx) Reset() {
	ctx.buf = ctx.buf[:0]
}
