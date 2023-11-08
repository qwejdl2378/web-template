package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qwejdl2378/web-template/controller"
	"k8s.io/klog/v2"
	"os"
)

func setUpRouter(engine *gin.Engine) {
	usrController := controller.NewUserController()
	engine.GET("/health", func(ctx *gin.Context) {
		ctx.AbortWithStatus(200)
	})
	v1 := engine.Group("v1")
	{
		v1.POST("user", usrController.CreateUser)
	}
}

func run() {
	addr := ":9099"
	port := os.Getenv("port")
	if port != "" {
		addr = fmt.Sprintf(":%s", port)
	}
	engine := gin.New()
	engine.Use(
		gin.Recovery(),
		gin.LoggerWithConfig(
			gin.LoggerConfig{
				SkipPaths: []string{"/health"},
			},
		),
	)
	setUpRouter(engine)
	if err := engine.Run(addr); err != nil {
		klog.Fatal(err)
	}
}

func Start() {
	run()
}
