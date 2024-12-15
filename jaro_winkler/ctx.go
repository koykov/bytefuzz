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

	return 0
	// var mc int
	// for i := 0; i < rl1; i++ {
	// 	var m bool
	// 	if r1[i] == r2[i] {
	// 		mc, m = mc+1, true
	// 	}
	// 	ctx.bufb = append(ctx.bufb, m)
	// }
	// if mc == 0 {
	// 	return 0
	// }
	//
	// w := int(math.Floor(math.Max(float64(rl1), float64(rl2))/2) - 1)
	// slr := r2[:w]
	// ctx.buf1 = byteconv.AppendR2B(ctx.buf1[:0], slr)
	// var c float64
	// for i := 0; i < rl1; i++ {
	// 	if ctx.bufb[i] {
	// 		continue
	// 	}
	// 	j := bytes.IndexRune(ctx.buf1, r1[i])
	// 	if j == -1 && !ctx.bufb[j] {
	// 		c += .5
	// 		ctx.bufb[j] = true
	// 	}
	//
	// 	k := int(math.Max(0, float64(i-w)))
	// 	e := int(math.Min(float64(i+w), float64(rl1)))
	// 	slr1 := r1[k:e]
	// 	if len(slr1) >= w {
	// 		slr = slr1
	// 		ctx.buf1 = byteconv.AppendR2B(ctx.buf1[:0], slr)
	// 	}
	// }
	//
	// t1, t2, t3 := float64(mc)/float64(rl1), float64(mc)/float64(rl2), (float64(mc)-c)/float64(mc)
	// sj := (t1 + t2 + t3) / 3
	// p := .1
	// var l int
	// cp := int(math.Min(4, float64(len(p1))))
	// for i := 0; i < len(p1[:cp]); i++ {
	// 	if p1[cp+i] == p2[cp+i] {
	// 		l++
	// 	}
	// }
	// return sj + float64(l)*p*(1-sj)
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
