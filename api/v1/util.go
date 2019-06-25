package v1

import (
	"log"

	"github.com/danclive/mqtt-console/db"
	"github.com/gin-gonic/gin"
)

func get_user(c *gin.Context) db.User {
	value, has := c.Get("user")
	if !has {
		log.Fatal("can't get user!")
	}

	user := value.(db.User)

	return user
}
