package api

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/danclive/mqtt-console/config"
	"github.com/danclive/mqtt-console/db"
	"github.com/danclive/mqtt-console/log"
	"go.mongodb.org/mongo-driver/bson"

	"time"

	"github.com/gin-gonic/gin"
)

var AuthMiddleware = GinJWTMiddleware{
	Realm:      "test zone",
	Key:        []byte("tstkey"),
	Timeout:    time.Hour * 24,
	MaxRefresh: time.Hour,
	Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
		encode_password := PassEncode(password)

		collection := db.MongoClient.Database(config.Config.Mongo.Database, nil).Collection("user", nil)

		result := collection.FindOne(context.Background(), bson.D{{"name", userId}}, &options.FindOneOptions{})
		if result.Err() != nil {
			log.Error(result.Err())
			return userId, false
		}

		var user db.User

		if err := result.Decode(&user); err != nil {
			log.WarnWithFields(err, log.Fields{"name": userId})
			return userId, false
		}

		if user.Pass == encode_password {
			return userId, true
		}

		log.WarnWithFields("密码错误", log.Fields{"name": userId})
		return userId, false
	},
	Authorizator: func(userId string, c *gin.Context) bool {

		collection := db.MongoClient.Database(config.Config.Mongo.Database, nil).Collection("user", nil)

		result := collection.FindOne(context.Background(), bson.D{{"name", userId}}, &options.FindOneOptions{})
		if result.Err() != nil {
			log.Error(result.Err())
			return false
		}

		var user db.User

		if err := result.Decode(&user); err != nil {
			log.WarnWithFields(err, log.Fields{"name": userId})
			return false
		}

		c.Set("user", user)

		return true
	},
	Unauthorized: func(c *gin.Context, code int, message string) {
		c.JSON(code, gin.H{
			"success": false,
			"message": gin.H{
				"code": code,
				"info": message,
			},
		})
	},
	TokenLookup:   "header:Authorization",
	TokenHeadName: "Bearer",
	TimeFunc:      time.Now,
}

func PassEncode(value string) (str string) {
	sha_ctx := sha256.New()
	sha_ctx.Write([]byte(value))

	cipher_str := sha_ctx.Sum(nil)

	str = hex.EncodeToString(cipher_str)
	return
}
