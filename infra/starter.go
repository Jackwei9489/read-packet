package infra

import "github.com/tietang/props/v3/kvs"

const (
	KeyProps = "props"
)

type StarterContext map[string]any

func (s StarterContext) Props() kvs.ConfigSource {
	p := s[KeyProps]
	if p == nil {
		panic("配置未被初始化")
	}
	return p.(kvs.ConfigSource)
}

type Starter interface {
	// Init 1.系统启动，初始化基础资源
	Init(StarterContext)
	// SetUp 2.系统基础资源安装
	SetUp(StarterContext)
	// Start 3. 启动基础资源
	Start(StarterContext)
	// StartBlocking 启动器是否可阻塞
	StartBlocking() bool
	// Stop 资源停止和销毁
	Stop(StarterContext)
}

var _ Starter = new(BaseStarter)

type BaseStarter struct {
}

func (b BaseStarter) Init(context StarterContext)  {}
func (b BaseStarter) SetUp(context StarterContext) {}
func (b BaseStarter) Start(context StarterContext) {}
func (b BaseStarter) StartBlocking() bool          { return false }
func (b BaseStarter) Stop(context StarterContext)  {}

type startRegister struct {
	starters []Starter
}

func (r *startRegister) Register(s Starter) {
	r.starters = append(r.starters, s)
}

func (r *startRegister) AllStarters() []Starter {
	return r.starters
}

var StartRegister = new(startRegister)

func Register(s Starter) {
	StartRegister.Register(s)
}
