package v1

import (
	"context"

	"github.com/danclive/mqtt-console/api/util"
	"github.com/danclive/mqtt-console/log"
	v1 "github.com/danclive/mqtt-console/model/v1"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func EmqxAuth(c *gin.Context) {
	var d struct {
		ClientId string `json:"clientid" form:"clientid" binding:"required"`
		Username string `json:"username" form:"username" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}

	if err := c.Bind(&d); err != nil {
		// c.JSON(util.Error(400, err.Error()))
		log.Error(err)
		c.Status(400)
		return
	}

	log.Debug(d)

	deviceColl := v1.GetDeviceColl()

	context := context.Background()

	result := deviceColl.FindOne(context,
		bson.M{"device_id": d.ClientId, "delete": false},
		&options.FindOneOptions{})

	if result.Err() != nil {
		// c.JSON(util.Error(500, result.Err().Error()))
		log.Error(result.Err())
		c.Status(500)
		return
	}

	var device v1.Device
	if err := result.Decode(&device); err != nil {
		if err.Error() == "mongo: no documents in result" {
			// c.JSON(util.Error(404, "404"))
			c.Status(404)
			return
		}

		log.Error(err)

		// c.JSON(util.Error(500, err.Error()))
		c.Status(500)
		return
	}

	userColl := v1.GetUserColl()

	result2 := userColl.FindOne(context, bson.M{"_id": device.UserID}, &options.FindOneOptions{})
	if result2.Err() != nil {
		// c.JSON(util.Error(500, result.Err().Error()))
		log.Error(result.Err())
		c.Status(500)
		return
	}

	var user v1.User
	if err := result2.Decode(&user); err != nil {
		if err.Error() == "mongo: no documents in result" {
			// c.JSON(util.Error(404, "404"))
			c.Status(404)
			return
		}

		log.Error(err)

		// c.JSON(util.Error(500, err.Error()))
		log.Error(result.Err())
		c.Status(500)
		return
	}

	if d.Username != user.Name || util.PassEncode(d.Password) != user.Pass {
		// c.JSON(util.Error(401, "401"))
		c.Status(401)
		return
	}

	c.Status(200)
}

func EmqxSuper(c *gin.Context) {
	var d struct {
		ClientId string `json:"clientid" form:"clientid" binding:"required"`
		Username string `json:"username" form:"username" binding:"required"`
	}

	if err := c.Bind(&d); err != nil {
		// c.JSON(util.Error(400, err.Error()))
		log.Error(err)
		c.Status(400)
		return
	}

	log.Debug(d)

	deviceColl := v1.GetDeviceColl()

	context := context.Background()

	result := deviceColl.FindOne(context,
		bson.M{"device_id": d.ClientId, "delete": false},
		&options.FindOneOptions{})

	if result.Err() != nil {
		// c.JSON(util.Error(500, result.Err().Error()))
		log.Error(result.Err())
		c.Status(500)
		return
	}

	var device v1.Device
	if err := result.Decode(&device); err != nil {
		if err.Error() == "mongo: no documents in result" {
			// c.JSON(util.Error(404, "404"))
			c.Status(404)
			return
		}

		log.Error(err)

		// c.JSON(util.Error(500, err.Error()))
		log.Error(err)
		c.Status(500)
		return
	}

	userColl := v1.GetUserColl()

	result2 := userColl.FindOne(context, bson.M{"_id": device.UserID}, &options.FindOneOptions{})
	if result2.Err() != nil {
		// c.JSON(util.Error(500, result.Err().Error()))
		log.Error(result.Err())
		c.Status(500)
		return
	}

	var user v1.User
	if err := result2.Decode(&user); err != nil {
		if err.Error() == "mongo: no documents in result" {
			// c.JSON(util.Error(404, "404"))
			c.Status(404)
			return
		}

		log.Error(err)

		// c.JSON(util.Error(500, err.Error()))
		log.Error(err)
		c.Status(500)
		return
	}

	if !user.Super {
		// c.JSON(util.Error(401, "401"))
		c.Status(401)
		return
	}

	c.Status(200)
}

func EmqxAcl(c *gin.Context) {
	var d struct {
		ClientId string `json:"clientid" form:"clientid" binding:"required"`
		Username string `json:"username" form:"username" binding:"required"`
		Access   string `json:"access" form:"access" binding:"required"`
		IpAddr   string `json:"ipaddr" form:"ipaddr" binding:"required"`
		Topic    string `json:"topic" form:"topic" binding:"required"`
	}

	if err := c.Bind(&d); err != nil {
		// c.JSON(util.Error(400, err.Error()))
		log.Error(err)
		c.Status(400)
		return
	}

	log.Debug(d)

	c.Status(200)

}
