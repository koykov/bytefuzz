package jaro_winkler

import "sync"

type pool struct {
	p sync.Pool
}

var p = pool{p: sync.Pool{New: func() interface{} { return &Ctx{} }}}

func Acquire() *Ctx {
	return p.p.Get().(*Ctx)
}

func Release(x *Ctx) {
	if x == nil {
		return
	}
	x.Reset()
	p.p.Put(x)
}

var _, _ = Acquire, Release
