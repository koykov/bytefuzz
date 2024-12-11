package jaro_winkler

import (
	"math"
	"unicode/utf8"

	"github.com/koykov/bytealg"
	"github.com/koykov/byteconv"
)

type Ctx struct {
	buf  []byte
	bufr []rune
	bufb []bool
}

func NewCtx() *Ctx {
	return &Ctx{}
}

func (ctx *Ctx) Distance(text, target []byte) float64 {
	text = bytealg.ToLowerBytes(text)
	target = bytealg.ToLowerBytes(target)
	return ctx.dist(text, target)
}

func (ctx *Ctx) DistanceString(text, target string) float64 {
	ctx.buf = append(ctx.buf, text...)
	ctx.buf = append(ctx.buf, target...)
	return ctx.Distance(ctx.buf[:len(text)], ctx.buf[len(text):])
}

func (ctx *Ctx) dist(p1, p2 []byte) float64 {
	rl1, rl2 := utf8.RuneCount(p1), utf8.RuneCount(p2)
	if rl1 > rl2 {
		p1, p2 = p2, p1
		rl1, rl2 = rl2, rl1
	}

	ctx.bufr = byteconv.AppendBytesToRunes(ctx.bufr, p1)
	ctx.bufr = byteconv.AppendBytesToRunes(ctx.bufr, p2)
	r1, r2 := ctx.bufr[:rl1], ctx.bufr[rl1:]

	var mc int
	for i := 0; i < rl1; i++ {
		var m bool
		if r1[i] == r2[i] {
			mc, m = mc+1, true
		}
		ctx.bufb = append(ctx.bufb, m)
	}
	if mc == 0 {
		return 0
	}

	w := int(math.Floor(math.Max(float64(rl1), float64(rl2))/2) - 1)
	_ = w
	// todo ...

	return 0
}

func (ctx *Ctx) Reset() {
	ctx.buf = ctx.buf[:0]
	ctx.bufr = ctx.bufr[:0]
	ctx.bufb = ctx.bufb[:0]
}
