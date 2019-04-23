package v1

import (
	"context"
	"time"

	"github.com/danclive/mqtt-console/api/util"
	"github.com/danclive/mqtt-console/config"
	"github.com/danclive/mqtt-console/db"
	"github.com/danclive/mqtt-console/log"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Userlist(c *gin.Context) {
	log.Info("user list")

	limit, offset, err := util.Page(c)

	delete := false
	if c.Query("delete") == "true" {
		delete = true
	}

	collection := db.MongoClient.Database(config.Config.Mongo.Database, nil).Collection("user", nil)

	context := context.Background()

	count, err := collection.CountDocuments(context, bson.M{"delete": delete}, &options.CountOptions{})
	if err != nil {
		c.JSON(util.Error(500, err.Error()))
		return
	}

	if count <= 0 {
		c.JSON(util.Error(404, "404"))
		return
	}

	cur, err := collection.Find(
		context,
		bson.M{"delete": delete},
		&options.FindOptions{Limit: &limit, Skip: &offset},
	)

	if err != nil {
		c.JSON(util.Error(500, err.Error()))
		return
	}

	defer cur.Close(context)

	arrays := make([]db.User, 0, 10)

	for cur.Next(context) {
		var user db.User
		log.Info(cur)
		if err := cur.Decode(&user); err != nil {
			c.JSON(util.Error(500, err.Error()))
			return
		}

		user.Pass = ""

		arrays = append(arrays, user)
	}

	c.JSON(util.Success(gin.H{
		"items": arrays,
		"count": count,
	}))
}

func UserDetail(c *gin.Context) {

	user_id := c.Param("id")

	if len(user_id) != 24 {
		c.JSON(util.Error(404, "404"))
		return
	}

	oid, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		c.JSON(util.Error(500, err.Error()))
		return
	}

	delete := false
	if c.Query("delete") == "true" {
		delete = true
	}

	collection := db.MongoClient.Database(config.Config.Mongo.Database, nil).Collection("user", nil)

	result := collection.FindOne(context.Background(), bson.M{"_id": oid, "delete": delete}, &options.FindOneOptions{})
	if result.Err() != nil {
		c.JSON(util.Error(500, result.Err().Error()))
		return
	}

	var user db.User
	if err = result.Decode(&user); err != nil {
		if err.Error() == "mongo: no documents in result" {
			c.JSON(util.Error(404, "404"))
			return
		}

		c.JSON(util.Error(500, err.Error()))
		return
	}

	c.JSON(util.Success(user))
}

func UserPost(c *gin.Context) {
	var u struct {
		Name  string `json:"name"`
		Pass  string `json:"pass"`
		Super bool   `json:"super"`
	}

	if err := c.Bind(&u); err != nil {
		c.JSON(util.Error(500, err.Error()))
		return
	}

	if u.Name == "" {
		c.JSON(util.Error(400, "name 不能为空"))
		return
	}

	if u.Pass == "" {
		c.JSON(util.Error(400, "pass 不能为空"))
		return
	}

	user := db.User{
		ID:    primitive.NewObjectID(),
		Name:  u.Name,
		Pass:  util.PassEncode(u.Pass),
		Time:  primitive.NewDateTimeFromTime(time.Now()),
		Super: u.Super,
	}

	log.Debug(user)

	doc, err := bson.Marshal(&user)
	if err != nil {
		c.JSON(util.Error(500, err.Error()))
		return
	}

	log.Debug(doc)

	collection := db.MongoClient.Database(config.Config.Mongo.Database, nil).Collection("user", nil)

	_, err = collection.InsertOne(context.Background(), doc, &options.InsertOneOptions{})
	if err != nil {
		c.JSON(util.Error(500, err.Error()))
		return
	}

	c.JSON(util.Success(gin.H{"id": user.ID}))
}

func UserPatch(c *gin.Context) {
	user_id := c.Param("id")

	if len(user_id) != 24 {
		c.JSON(util.Error(404, "404"))
		return
	}

	oid, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		c.JSON(util.Error(500, err.Error()))
		return
	}

	delete := false
	if c.Query("delete") == "true" {
		delete = true
	}

	collection := db.MongoClient.Database(config.Config.Mongo.Database, nil).Collection("user", nil)

	result := collection.FindOne(context.Background(), bson.M{"_id": oid, "delete": delete}, &options.FindOneOptions{})
	if result.Err() != nil {
		c.JSON(util.Error(500, result.Err().Error()))
		return
	}

	var user db.User
	if err = result.Decode(&user); err != nil {
		if err.Error() == "mongo: no documents in result" {
			c.JSON(util.Error(404, "404"))
			return
		}

		c.JSON(util.Error(500, err.Error()))
		return
	}

	var u struct {
		Name  string `json:"name"`
		Pass  string `json:"pass"`
		Super string `json:"super"`
	}

	if err := c.Bind(&u); err != nil {
		c.JSON(util.Error(500, err.Error()))
		return
	}

	doc := bson.M{}

	if u.Name != "" {
		doc["name"] = u.Name
	}

	if u.Pass != "" {
		doc["pass"] = util.PassEncode(u.Pass)
	}

	if u.Super == "true" {
		doc["super"] = true
	} else if u.Super == "false" {
		doc["super"] = false
	}

	_, err = collection.UpdateOne(context.Background(), bson.M{"_id": oid}, bson.M{"$set": doc}, &options.UpdateOptions{})
	if err != nil {
		c.JSON(util.Error(500, err.Error()))
		return
	}

	c.JSON(util.Success(gin.H{"id": user_id}))
}

func UserDelete(c *gin.Context) {
	user_id := c.Param("id")

	if len(user_id) != 24 {
		c.JSON(util.Error(404, "404"))
		return
	}

	oid, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		c.JSON(util.Error(500, err.Error()))
		return
	}

	collection := db.MongoClient.Database(config.Config.Mongo.Database, nil).Collection("user", nil)

	result := collection.FindOne(context.Background(), bson.M{"_id": oid, "delete": false}, &options.FindOneOptions{})
	if result.Err() != nil {
		c.JSON(util.Error(500, result.Err().Error()))
		return
	}

	var user db.User
	if err = result.Decode(&user); err != nil {
		if err.Error() == "mongo: no documents in result" {
			c.JSON(util.Error(404, "404"))
			return
		}

		c.JSON(util.Error(500, err.Error()))
		return
	}

	_, err = collection.UpdateOne(context.Background(), bson.M{"_id": oid}, bson.M{"$set": bson.M{"delete": true}}, &options.UpdateOptions{})
	if err != nil {
		c.JSON(util.Error(500, err.Error()))
		return
	}

	c.JSON(util.Success(gin.H{"id": user_id}))
}

func UserReset(c *gin.Context) {
	user_id := c.Param("id")

	if len(user_id) != 24 {
		c.JSON(util.Error(404, "404"))
		return
	}

	oid, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		c.JSON(util.Error(500, err.Error()))
		return
	}

	collection := db.MongoClient.Database(config.Config.Mongo.Database, nil).Collection("user", nil)

	result := collection.FindOne(context.Background(), bson.M{"_id": oid, "delete": true}, &options.FindOneOptions{})
	if result.Err() != nil {
		c.JSON(util.Error(500, result.Err().Error()))
		return
	}

	var user db.User
	if err = result.Decode(&user); err != nil {
		if err.Error() == "mongo: no documents in result" {
			c.JSON(util.Error(404, "404"))
			return
		}

		c.JSON(util.Error(500, err.Error()))
		return
	}

	_, err = collection.UpdateOne(context.Background(), bson.M{"_id": oid}, bson.M{"$set": bson.M{"delete": false}}, &options.UpdateOptions{})
	if err != nil {
		c.JSON(util.Error(500, err.Error()))
		return
	}

	c.JSON(util.Success(gin.H{"id": user_id}))
}
