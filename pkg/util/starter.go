package util

const (
	KeyProps = "_conf"
	// DefaultPriority 默认优先级
	DefaultPriority = 10000
	// SystemGroup 系统优先级
	SystemGroup PriorityGroup = 30
	// BasicResourcesGroup 基础资源优先级
	BasicResourcesGroup PriorityGroup = 20
	// AppGroup app优先级
	AppGroup PriorityGroup = 10
	// IntMax
	IntMax = int(^uint(0) >> 1)
)

type (
	// StarterContext 资源启动器上下文，
	// 用来在服务资源初始化、安装、启动和停止的生命周期中变量和对象的传递
	StarterContext map[string]interface{}
	// PriorityGroup 优先级
	PriorityGroup int
	// Starters 所有资源启动器
	Starters []Starter
)

// Starter 资源启动器
type Starter interface {
	//资源初始化和，通常把一些准备资源放在这里运行
	Init(StarterContext)
	//资源的安装，所有启动需要的具备条件，使得资源达到可以启动的就备状态
	Setup(StarterContext)
	//启动资源，达到可以使用的状态
	Start(StarterContext)
	//说明该资源启动器开始启动服务时，是否会阻塞
	//如果存在多个阻塞启动器时，只有最后一个阻塞，之前的会通过goroutine来异步启动
	//所以，需要规划好启动器注册顺序
	StartBlocking() bool
	//资源停止：
	// 通常在启动时遇到异常时或者启用远程管理时，用于释放资源和终止资源的使用，
	// 通常要优雅的释放，等待正在进行的任务继续，但不再接受新的任务
	Stop(StarterContext)
	PriorityGroup() PriorityGroup
	Priority() int
}

// BaseStarter 默认的空实现,方便资源启动器的实现
type BaseStarter struct{}

func (s *BaseStarter) Init(ctx StarterContext)      {}
func (s *BaseStarter) Setup(ctx StarterContext)     {}
func (s *BaseStarter) Start(ctx StarterContext)     {}
func (s *BaseStarter) Stop(ctx StarterContext)      {}
func (s *BaseStarter) StartBlocking() bool          { return false }
func (s *BaseStarter) PriorityGroup() PriorityGroup { return BasicResourcesGroup }
func (s *BaseStarter) Priority() int                { return DefaultPriority }

type ConfigSource interface {
	Name() string
}

func (s StarterContext) Props() ConfigSource {
	p := s[KeyProps]
	if p == nil {
		panic("配置还没有被初始化")
	}
	return p.(ConfigSource)
}

func (s StarterContext) SetProps(conf ConfigSource) {
	s[KeyProps] = conf
}

// Len 启动器数量
func (s Starters) Len() int { return len(s) }

// Swap 交换
func (s Starters) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Less 判断优先级
func (s Starters) Less(i, j int) bool {
	return s[i].PriorityGroup() > s[j].PriorityGroup() && s[i].Priority() > s[j].Priority()
}

// starterRegister 服务启动注册器
// 全局只有一个
type starterRegister struct {
	nonBlockingStarters []Starter
	blockingStarters    []Starter
}
