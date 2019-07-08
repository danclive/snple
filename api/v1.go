package api

import (
	"net/http"

	"github.com/danclive/mqtt-console/api/util"
	v1 "github.com/danclive/mqtt-console/api/v1"
	"github.com/danclive/mqtt-console/log"
	modelv1 "github.com/danclive/mqtt-console/model/v1"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func V1(engine gin.IRouter) {
	login := engine.Group("")
	login.POST("/login", AuthMiddleware.LoginHandler)

	user := engine.Group("user", AuthMiddleware.MiddlewareFunc())
	user.GET("/", super, v1.UserList)
	user.GET("/:id", super, v1.UserDetail)
	user.POST("/", super, v1.UserPost)
	user.PATCH("/:id", super, v1.UserPatch)
	user.DELETE("/:id", super, v1.UserDelete)
	user.DELETE("/:id/reset", super, v1.UserReset)

	mine := engine.Group("mine", AuthMiddleware.MiddlewareFunc())
	mine.GET("/", v1.Mine)

	user.GET("/:id/device", super, v1.UserDeviceList)
	user.POST("/:id/device", super, v1.UserDevicePost)

	device := engine.Group("device", AuthMiddleware.MiddlewareFunc())
	device.GET("/", v1.Devicelist)
	device.GET("/:id", v1.DeviceDetail)
	device.POST("/", v1.DevicePost)
	device.PATCH("/:id", v1.DevicePatch)
	device.DELETE("/:id", v1.DeviceDelete)
	device.DELETE("/:id/reset", v1.DeviceReset)

	emqx := engine.Group("emqx")
	emqx.POST("/auth", v1.EmqxAuth)
	emqx.GET("/super", v1.EmqxSuper)
	emqx.POST("/acl", v1.EmqxAcl)

	gutil := engine.Group("util")
	gutil.GET("/genid", func(c *gin.Context) {
		c.JSON(util.Success(gin.H{"id": primitive.NewObjectID()}))
	})
}

func super(c *gin.Context) {
	json := gin.H{
		"success": false,
		"message": gin.H{
			"code": 403,
			"info": "权限不足: 你不是管理员!",
		},
	}

	value, has := c.Get("user")
	if !has {
		c.AbortWithStatusJSON(http.StatusForbidden, json)
	}

	user := value.(modelv1.User)

	if !user.Super {
		log.WarnWithFields("super: ", log.Fields{"user": user})
		c.AbortWithStatusJSON(http.StatusForbidden, json)
	}
}
