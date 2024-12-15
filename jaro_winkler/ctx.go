package jaro_winkler

import (
	"github.com/koykov/bytealg"
	"github.com/koykov/byteconv"
)

type Ctx struct {
	buf  []byte
	buf1 []byte
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
	ctx.bufr = byteconv.AppendBytesToRunes(ctx.bufr, p1)
	r1, rl1 := ctx.bufr, len(ctx.bufr)
	ctx.bufr = byteconv.AppendBytesToRunes(ctx.bufr, p2)
	r2, rl2 := ctx.bufr[rl1:], len(ctx.bufr)-rl1

	if rl1 == 0 || rl2 == 0 {
		return 0
	}

	var ml, sr int
	if ml = rl1; ml < rl2 {
		ml = rl2
	}
	if sr = (ml / 2) - 1; sr < 0 {
		sr = 0
	}

	ctx.bufb = growBool(ctx.bufb, rl1+rl2)
	f1, f2 := ctx.bufb[:rl1], ctx.bufb[:rl1]

	var cc int
	for i := 0; i < rl1; i++ {
		lo := 0
		if i > sr {
			lo = i - sr
		}
		hi := rl2 - 1
		if i+sr < rl2 {
			hi = i + sr
		}
		for j := lo; j < hi; j++ {
			if !ctx.bufb[j] && r2[j] == r1[i] {
				f1[i] = true
				f2[j] = true
				cc++
				break
			}
		}
	}
	if cc == 0 {
		return 0
	}

	var tc, k int
	for i := 0; i < rl1; i++ {
		if f1[i] {
			j := k
			for ; j < rl2; j++ {
				if f2[j] {
					k = j + 1
					break
				}
			}
			if r1[i] != r2[j] {
				tc++
			}
		}
	}
	tc /= 2

	w := (float64(cc)/float64(rl1) + float64(cc)/float64(rl2) + (float64(cc)-float64(tc))/float64(cc)) / 3
	if w > .7 && rl1 > 3 && rl2 > 3 {
		i, j := 0, 4
		if j < ml {
			j = ml
		}
		for i < j && r1[i] == r2[i] {
			i++
		}
		if i > 0 {
			w += float64(i) * .1 * (1 - w)
		}
	}
	return w
}

func (ctx *Ctx) Reset() {
	ctx.buf = ctx.buf[:0]
	ctx.buf1 = ctx.buf1[:0]
	ctx.bufr = ctx.bufr[:0]
	ctx.bufb = ctx.bufb[:0]
}

func growBool(buf []bool, ln int) []bool {
	if cap(buf) >= ln {
		return buf[:ln]
	}
	buf = append(buf, make([]bool, cap(buf)+ln)...)
	return buf[:ln]
}
