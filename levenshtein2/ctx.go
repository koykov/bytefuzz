package levenshtein2

import (
	"math"

	"github.com/koykov/byteconv"
)

type Ctx struct {
	text, target []rune

	buf     [math.MaxUint16]uint64
	bufHSZ2 []uint64
}

func NewCtx() *Ctx {
	return &Ctx{}
}

func (ctx *Ctx) Distance(text, target []byte) float64 {
	a, b := text, target
	if len(b) == 0 {
		return float64(len(a))
	}
	if len(a) == 0 {
		return float64(len(b))
	}
	if len(a) < len(b) {
		a, b = b, a
	}

	ctx.text = byteconv.AppendB2R(ctx.text[:0], a)
	ctx.target = byteconv.AppendB2R(ctx.target[:0], b)
	if len(a) <= 64 {
		return ctx.dist64(ctx.text, ctx.target)
	}
	return ctx.distN(ctx.text, ctx.target)
}

func (ctx *Ctx) DistanceString(text, target string) float64 {
	a, b := text, target
	if len(b) == 0 {
		return float64(len(a))
	}
	if len(a) == 0 {
		return float64(len(b))
	}
	if len(a) < len(b) {
		a, b = b, a
	}

	ctx.text = byteconv.AppendS2R(ctx.text[:0], a)
	ctx.target = byteconv.AppendS2R(ctx.target[:0], b)
	if len(a) <= 64 {
		return ctx.dist64(ctx.text, ctx.target)
	}
	return ctx.distN(ctx.text, ctx.target)
}

func (ctx *Ctx) dist64(a, b []rune) (sc float64) {
	pv := ^uint64(0)
	mv := uint64(0)
	for i := 0; i < len(a); i++ {
		ctx.buf[a[i]] |= uint64(1) << uint64(sc)
		sc++
	}
	ls := uint64(1) << uint64(sc-1)
	_ = b[len(b)-1]
	for i := 0; i < len(b); i++ {
		eq := ctx.buf[b[i]]
		xv := eq | mv
		eq |= ((eq & pv) + pv) ^ pv
		mv |= ^(eq | pv)
		pv &= eq
		if (mv & ls) != 0 {
			sc++
		}
		if (pv & ls) != 0 {
			sc--
		}
		mv = (mv << 1) | 1
		pv = (pv << 1) | ^(xv | mv)
		mv &= xv
	}
	_ = a[len(a)-1]
	for i := 0; i < len(a); i++ {
		ctx.buf[a[i]] = 0
	}
	return sc
}

func (ctx *Ctx) distN(s1, s2 []rune) float64 {
	n := len(s1)
	m := len(s2)
	_, _ = s1[n-1], s2[n-1]

	hsize := ((n - 1) / 64) + 1
	vsize := ((m - 1) / 64) + 1
	if cap(ctx.bufHSZ2) < hsize*2 {
		ctx.bufHSZ2 = make([]uint64, 0, hsize*2)
	}
	ctx.bufHSZ2 = ctx.bufHSZ2[:hsize*2]
	phc, mhc := ctx.bufHSZ2[:hsize], ctx.bufHSZ2[hsize:]
	for i := 0; i < hsize; i++ {
		phc[i] = ^uint64(0)
		mhc[i] = 0
	}

	var j int
	for ; j < vsize-1; j++ {
		mv := uint64(0)
		pv := ^uint64(0)
		start := j * 64
		vlen := mn(64, m) + start
		for k := start; k < vlen; k++ {
			ctx.buf[s2[k]] |= uint64(1) << (k & 63)
		}

		for i := 0; i < n; i++ {
			eq := ctx.buf[s1[i]]
			pb := (phc[i/64] >> (i & 63)) & 1
			mb := (mhc[i/64] >> (i & 63)) & 1
			xv := eq | mv
			xh := ((((eq | mb) & pv) + pv) ^ pv) | eq | mb
			ph := mv | ^(xh | pv)
			mh := pv & xh
			if ((ph >> 63) ^ pb) != 0 {
				phc[i/64] ^= uint64(1) << (i & 63)
			}
			if ((mh >> 63) ^ mb) != 0 {
				mhc[i/64] ^= uint64(1) << (i & 63)
			}
			ph = (ph << 1) | pb
			mh = (mh << 1) | mb
			pv = mh | ^(xv | ph)
			mv = ph & xv
		}
		for k := start; k < vlen; k++ {
			ctx.buf[s2[k]] = 0
		}
	}
	mv := uint64(0)
	pv := ^uint64(0)
	start := j * 64
	vlen := mn(64, m-start) + start
	for k := start; k < vlen; k++ {
		ctx.buf[s2[k]] |= uint64(1) << (k & 63)
	}
	sc := uint64(m)
	for i := 0; i < n; i++ {
		eq := ctx.buf[s1[i]]
		pb := (phc[i/64] >> (i & 63)) & 1
		mb := (mhc[i/64] >> (i & 63)) & 1
		xv := eq | mv
		xh := ((((eq | mb) & pv) + pv) ^ pv) | eq | mb
		ph := mv | ^(xh | pv)
		mh := pv & xh
		sc += (ph >> ((m - 1) & 63)) & 1
		sc -= (mh >> ((m - 1) & 63)) & 1
		if ((ph >> 63) ^ pb) != 0 {
			phc[i/64] ^= uint64(1) << (i & 63)
		}
		if ((mh >> 63) ^ mb) != 0 {
			mhc[i/64] ^= uint64(1) << (i & 63)
		}
		ph = (ph << 1) | pb
		mh = (mh << 1) | mb
		pv = mh | ^(xv | ph)
		mv = ph & xv
	}
	for k := start; k < vlen; k++ {
		ctx.buf[s2[k]] = 0
	}
	return float64(sc)
}

func (ctx *Ctx) Reset() {
	ctx.text = ctx.text[:0]
	ctx.target = ctx.target[:0]
	ctx.bufHSZ2 = ctx.bufHSZ2[:0]
}

func mn(a, b int) int {
	if a < b {
		return a
	}
	return b
}
