package v1

import (
	"log"

	v1 "github.com/danclive/mqtt-console/model/v1"
	"github.com/gin-gonic/gin"
)

func get_user(c *gin.Context) v1.User {
	value, has := c.Get("user")
	if !has {
		log.Fatal("can't get user!")
	}

	user := value.(v1.User)

	return user
}
