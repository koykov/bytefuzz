package levenshtein

import (
	"sync"

	"github.com/koykov/bytefuzz"
)

type pool struct {
	p sync.Pool
}

var p = pool{p: sync.Pool{New: func() interface{} { return &ctx{} }}}

func Acquire() bytefuzz.Interface {
	return p.p.Get().(*ctx)
}

func Release(x bytefuzz.Interface) {
	if x == nil {
		return
	}
	x.Reset()
	p.p.Put(x)
}

var _, _ = Acquire, Release
