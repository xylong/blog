package base

import (
	initial "blog/init"
	"github.com/tietang/props/kvs"
)

var props kvs.ConfigSource

func Props() kvs.ConfigSource {
	return props
}

// PropsStarter 配置
type PropsStarter struct {
	initial.BaseStarter
}

func (p *PropsStarter) Init(ctx initial.StarterContext) {
	props = ctx.Props()
}
