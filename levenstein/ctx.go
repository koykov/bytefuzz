package levenstein

type Ctx struct {
	text, target []rune

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
	return 0
}

func (ctx *Ctx) DistanceString(text, target string) int {
	return 0
}

func (ctx *Ctx) Reset() {
	ctx.text = ctx.text[:0]
	ctx.target = ctx.target[:0]
	ctx.costIns = 0
	ctx.costDel = 0
	ctx.costRepl = 0
	ctx.matchFn = nil
}
