package levenstein

import "sync"

type p struct {
	p sync.Pool
}

var p_ p

func Acquire() *Ctx {
	x := p_.p.Get().(*Ctx)
	if x == nil {
		x = &Ctx{}
	}
	return x
}

func Release(x *Ctx) {
	if x == nil {
		return
	}
	p_.p.Put(x)
}
