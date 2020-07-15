package gojieba

import (
	"github.com/xu215740578/gojieba/deps/cppjieba"
	"github.com/xu215740578/gojieba/deps/limonp"
	"github.com/xu215740578/gojieba/dict"
)

func init() {
	dict.Init()
	limonp.Init()
	cppjieba.Init()
}
