package v1

import (
	"context"
	"time"

	"github.com/danclive/mqtt-console/api/util"
	"github.com/danclive/mqtt-console/log"
	v1 "github.com/danclive/mqtt-console/model/v1"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Mine(c *gin.Context) {
	value, has := c.Get("user")
	if !has {
		c.JSON(util.Error(404, "404"))
		return
	}

	user := value.(v1.User)

	c.JSON(util.Success(user))
}

func UserList(c *gin.Context) {
	limit, offset, err := util.Page(c)
	if err != nil {
		c.JSON(util.Error(400, "分页参数错误"))
		return
	}

	delete := false
	if c.Query("delete") == "true" {
		delete = true
	}

	collection := v1.GetUserColl()

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

	arrays := make([]v1.User, 0, 10)

	for cur.Next(context) {
		var user v1.User
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

	collection := v1.GetUserColl()

	result := collection.FindOne(context.Background(), bson.M{"_id": oid, "delete": delete}, &options.FindOneOptions{})
	if result.Err() != nil {
		c.JSON(util.Error(500, result.Err().Error()))
		return
	}

	var user v1.User
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
		Desc  string `json:"desc"`
	}

	if err := c.Bind(&u); err != nil {
		c.JSON(util.Error(400, err.Error()))
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

	collection := v1.GetUserColl()

	count, err := collection.CountDocuments(context.Background(), bson.M{"name": u.Name}, &options.CountOptions{})
	if err != nil {
		c.JSON(util.Error(500, err.Error()))
		return
	}

	if count > 0 {
		c.JSON(util.Error(409, "用户名已存在"))
		return
	}

	user := v1.User{
		ID:    primitive.NewObjectID(),
		Name:  u.Name,
		Pass:  util.PassEncode(u.Pass),
		Time:  time.Now(),
		Super: u.Super,
		Desc:  u.Desc,
	}

	log.Debug(user)

	doc, err := bson.Marshal(&user)
	if err != nil {
		c.JSON(util.Error(500, err.Error()))
		return
	}

	log.Debug(doc)

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

	collection := v1.GetUserColl()

	result := collection.FindOne(context.Background(), bson.M{"_id": oid, "delete": delete}, &options.FindOneOptions{})
	if result.Err() != nil {
		c.JSON(util.Error(500, result.Err().Error()))
		return
	}

	var user v1.User
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
		Desc  string `json:"Desc"`
	}

	if err := c.Bind(&u); err != nil {
		c.JSON(util.Error(400, err.Error()))
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

	if u.Desc != "" {
		doc["desc"] = u.Desc
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

	collection := v1.GetUserColl()

	result := collection.FindOne(context.Background(), bson.M{"_id": oid, "delete": false}, &options.FindOneOptions{})
	if result.Err() != nil {
		c.JSON(util.Error(500, result.Err().Error()))
		return
	}

	var user v1.User
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

	collection := v1.GetUserColl()

	result := collection.FindOne(context.Background(), bson.M{"_id": oid, "delete": true}, &options.FindOneOptions{})
	if result.Err() != nil {
		c.JSON(util.Error(500, result.Err().Error()))
		return
	}

	var user v1.User
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
