package base

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"red-packet/apis"
	"red-packet/infra"
	"time"
)

var engine *gin.Engine

func GinEngine() *gin.Engine {
	return engine
}

type GinServerStarter struct {
	infra.BaseStarter
}

func (s *GinServerStarter) Init(ctx infra.StarterContext) {
	// 创建gin实例
	engine = initGin()
	// 日志组件配置和扩展
}

func initGin() *gin.Engine {
	app := gin.Default()
	//app.Use(gin.Recovery())
	//app.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
	//	return fmt.Sprintf("| %s | %s | %d | %s  | %s |  %s |  %s |",
	//		params.TimeStamp.Format("2006-01-02.15:04:05.000000"),
	//		params.Latency.String(),
	//		params.StatusCode,
	//		params.ClientIP,
	//		params.Method,
	//		params.Path,
	//		params.ErrorMessage,
	//	)
	//}))
	app.Use(ginzap.Ginzap(logger.Desugar(), time.RFC3339, true))
	app.Use(ginzap.RecoveryWithZap(logger.Desugar(), true))
	return app
}

func (s *GinServerStarter) Start(ctx infra.StarterContext) {
	// 注册route
	registerRoutes()
	routes := engine.Routes()
	for _, v := range routes {
		logger.Infof("register route: %s:%s\n", v.Method, v.Path)
	}
	port := ctx.Props().GetDefault("app.server.port", "1234")

	err := engine.Run(":" + port)
	if err != nil {
		panic(err)
	}
}

func registerRoutes() {
	engine.GET("/", apis.Index)
}

func (s *GinServerStarter) StartBlocking() bool {
	return true
}
