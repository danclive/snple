package api

import (
	"net/http"

	v1 "github.com/danclive/mqtt-console/api/v1"
	"github.com/danclive/mqtt-console/db"
	"github.com/danclive/mqtt-console/log"
	"github.com/gin-gonic/gin"
)

func V1(engine gin.IRouter) {
	login := engine.Group("user")
	login.POST("/login", AuthMiddleware.LoginHandler)

	user := engine.Group("user", AuthMiddleware.MiddlewareFunc())
	user.GET("/", super, v1.Userlist)
	user.GET("/:id", super, v1.UserDetail)
	user.POST("/", super, v1.UserPost)
	user.PATCH("/:id", super, v1.UserPatch)
	user.DELETE("/:id", super, v1.UserDelete)
	user.DELETE("/:id/reset", super, v1.UserReset)
}

func super(c *gin.Context) {
	json := gin.H{
		"success": false,
		"message": gin.H{
			"code": 403,
			"info": "权限不足: 您不是管理员!",
		},
	}

	value, has := c.Get("user")
	if !has {
		c.AbortWithStatusJSON(http.StatusForbidden, json)
	}

	user := value.(db.User)

	if !user.Super {
		log.WarnWithFields("super: ", log.Fields{"user": user})
		c.AbortWithStatusJSON(http.StatusForbidden, json)
	}
}
