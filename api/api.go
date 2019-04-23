package api

import (
	"github.com/danclive/mqtt-console/log"

	"github.com/danclive/mqtt-console/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitApi() {

}

func RunApi() {
	log.Info("启动 web 服务 ...")
	if !config.Config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.Default()

	cors_config := cors.DefaultConfig()
	cors_config.AllowAllOrigins = true
	cors_config.AddAllowHeaders("Authorization")
	cors_config.AddAllowHeaders("Token")
	cors_config.AddAllowMethods("OPTIONS")

	app.Use(cors.New(cors_config))

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	V1(app.Group("/v1"))

	log.Info("web 服务运行在: ", config.Config.Api.Addr)
	err := app.Run(config.Config.Api.Addr) // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Fatal("web 服务启动失败: ", err)
	}

	log.Info("OK!")
}
