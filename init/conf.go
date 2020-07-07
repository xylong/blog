package init

import (
	"blog/pkg/util"
	"fmt"
)

func init() {
	util.Register(&ConfStarter{})
}

type ConfStarter struct {
	util.BaseStarter
}

func (c *ConfStarter) Init(ctx util.StarterContext) {
	fmt.Println("配置初始化")
}

func (c *ConfStarter) Setup(ctx util.StarterContext) {
	fmt.Println("配置安装")
}

func (c *ConfStarter) Start(ctx util.StarterContext) {
	fmt.Println("配置启动")
}
