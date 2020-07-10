package init

import "github.com/tietang/props/kvs"

// NewBoot 实例化启动管理器
func NewBoot(conf kvs.ConfigSource) *BootApplication {
	b := &BootApplication{
		conf:           conf,
		starterContext: StarterContext{},
	}
	b.starterContext[KeyProps] = conf
	return b
}

// BootApplication 应用程序启动管理器
type BootApplication struct {
	conf           kvs.ConfigSource
	starterContext StarterContext
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
	for _, starter := range StarterRegister.AllStarters() {
		starter.Init(b.starterContext)
	}
}

func (b *BootApplication) setup() {
	for _, starter := range StarterRegister.AllStarters() {
		starter.Setup(b.starterContext)
	}
}

func (b *BootApplication) start() {
	starters := StarterRegister.AllStarters()
	number := len(starters)
	for index, starter := range starters {
		if starter.StartBlocking() {
			// 如果是最后一个可阻塞的，直接启动并阻塞
			if index+1 == number {
				starter.Start(b.starterContext)
			} else {
				// 如果不是最后一个，则使用goroutine异步启动，防止阻塞后面的starter
				go starter.Start(b.starterContext)
			}
		} else {
			starter.Start(b.starterContext)
		}
	}
}
