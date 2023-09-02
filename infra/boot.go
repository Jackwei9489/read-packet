package infra

import "github.com/tietang/props/v3/kvs"

type BootApplication struct {
	conf           kvs.ConfigSource
	starterContext StarterContext
}

func New(conf kvs.ConfigSource) *BootApplication {
	b := &BootApplication{
		conf:           conf,
		starterContext: StarterContext{},
	}
	b.starterContext[KeyProps] = conf
	return b
}

func (b *BootApplication) Start() {
	// 1.初始化
	b.init()
	// 2. 装载
	b.setUp()
	// 3. 启动
	b.start()
}

func (b *BootApplication) init() {
	for _, starter := range StartRegister.AllStarters() {
		starter.Init(b.starterContext)
	}
}

func (b *BootApplication) setUp() {
	for _, starter := range StartRegister.AllStarters() {
		starter.SetUp(b.starterContext)
	}
}

func (b *BootApplication) start() {
	for i, starter := range StartRegister.AllStarters() {
		if starter.StartBlocking() {
			if i+1 == len(StartRegister.AllStarters()) {
				starter.Start(b.starterContext)
			} else {
				go starter.Start(b.starterContext)
			}
		} else {
			starter.Start(b.starterContext)
		}
	}
}
