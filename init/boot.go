package init

import "blog/pkg/util"

// NewBoot 实例化启动管理器
func NewBoot(conf interface{}) *BootApplication {
	return &BootApplication{
		conf:         conf,
		startContext: util.StarterContext{},
	}
}

// BootApplication 应用程序启动管理器
type BootApplication struct {
	conf         interface{}
	startContext util.StarterContext
}

// Start 启动所有启动器
func (b *BootApplication) Start() {
	// 1.初始化所有starter
	b.init()
	// 2.安装所有starter
	b.setup()
	// 3.启动所有starter
	b.start()
}

func (b *BootApplication) init() {
	for _, starter := range util.StarterRegister.AllStarters() {
		starter.Init(b.startContext)
	}
}

func (b *BootApplication) setup() {
	for _, starter := range util.StarterRegister.AllStarters() {
		starter.Setup(b.startContext)
	}
}

func (b *BootApplication) start() {
	starters := util.StarterRegister.AllStarters()
	number := len(starters)
	for index, starter := range starters {
		if starter.StartBlocking() {
			// 如果是最后一个可阻塞的，直接启动并阻塞
			if index+1 == number {
				starter.Start(b.startContext)
			} else {
				// 如果不是最后一个，则使用goroutine异步启动，防止阻塞后面的starter
				go starter.Start(b.startContext)
			}
		} else {
			starter.Start(b.startContext)
		}
	}
}
