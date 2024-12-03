package levenstein

import "github.com/koykov/byteconv"

type Ctx struct {
	text, target []rune
	m            [][]int

	costIns, costDel, costRepl int

	matchFn func(a, b rune) bool
}

func NewCtx(costIns, costDel, costRepl int, matchFn func(a, b rune) bool) *Ctx {
	return &Ctx{
		costIns:  costIns,
		costDel:  costDel,
		costRepl: costRepl,
		matchFn:  matchFn,
	}
}

func (ctx *Ctx) Distance(text, target []byte) int {
	ctx.text = byteconv.AppendB2R(ctx.text[:0], text)
	ctx.target = byteconv.AppendB2R(ctx.target[:0], target)
	return ctx.dist(ctx.text, ctx.target)
}

func (ctx *Ctx) DistanceString(text, target string) int {
	ctx.text = byteconv.AppendS2R(ctx.text[:0], text)
	ctx.target = byteconv.AppendS2R(ctx.target[:0], target)
	return ctx.dist(ctx.text, ctx.target)
}

func (ctx *Ctx) dist(text, target []rune) int {
	w, h := len(text)+1, len(target)+1
	for i := 0; i < 2; i++ {
		if len(ctx.m) < cap(ctx.m) {
			ctx.m = ctx.m[:i+1]
		} else {
			ctx.m = append(ctx.m, nil)
		}
		ctx.m[i] = append(ctx.m[i], i*ctx.costDel)
	}
	for i := 1; i < w; i++ {
		ctx.m[0] = append(ctx.m[0], i*ctx.costIns)
	}

	for i := 1; i < h; i++ {
		c, p := ctx.m[i%2], ctx.m[(i-1)%2]
		c[0] = i * ctx.costDel
		for j := 1; j < w; j++ {
			dc := p[j] + ctx.costDel
			sc := p[j-1]
			if !ctx.matchFn(ctx.text[i-1], ctx.target[j-1]) {
				sc += ctx.costRepl
			}
			ic := c[j-1] + ctx.costIns
			c[j] = min3(dc, sc, ic)
		}
	}
	return ctx.m[(h-1)%2][w-1]
}

func (ctx *Ctx) Reset() {
	ctx.m = ctx.m[:0]
	ctx.text = ctx.text[:0]
	ctx.target = ctx.target[:0]
	ctx.costIns = 0
	ctx.costDel = 0
	ctx.costRepl = 0
	ctx.matchFn = nil
}

func min3(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	}
	if b < c {
		return b
	}
	return c
}
