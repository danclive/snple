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

func UserDeviceList(c *gin.Context) {
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

	limit, offset, err := util.Page(c)
	if err != nil {
		c.JSON(util.Error(400, "分页参数错误"))
		return
	}

	delete := false
	if c.Query("delete") == "true" {
		delete = true
	}

	collection := v1.GetDeviceColl()

	context := context.Background()

	count, err := collection.CountDocuments(context, bson.M{"user_id": oid, "delete": delete}, &options.CountOptions{})
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
		bson.M{"user_id": oid, "delete": delete},
		&options.FindOptions{Limit: &limit, Skip: &offset},
	)

	if err != nil {
		c.JSON(util.Error(500, err.Error()))
		return
	}

	defer cur.Close(context)

	arrays := make([]v1.Device, 0, 10)

	for cur.Next(context) {
		var device v1.Device
		if err := cur.Decode(&device); err != nil {
			c.JSON(util.Error(500, err.Error()))
			return
		}

		arrays = append(arrays, device)
	}

	c.JSON(util.Success(gin.H{
		"items": arrays,
		"count": count,
	}))
}

func UserDevicePost(c *gin.Context) {
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

	var d struct {
		DeviceId string `json:"device_id"`
		Desc     string `json:"desc"`
	}

	if err := c.Bind(&d); err != nil {
		c.JSON(util.Error(400, err.Error()))
		return
	}

	if d.DeviceId == "" {
		c.JSON(util.Error(400, "device_id 不能为空"))
		return
	}

	collection := v1.GetDeviceColl()

	count, err := collection.CountDocuments(context.Background(), bson.M{"device_id": d.DeviceId}, &options.CountOptions{})
	if err != nil {
		c.JSON(util.Error(500, err.Error()))
		return
	}

	if count > 0 {
		c.JSON(util.Error(400, "device_id 重复"))
		return
	}

	device := v1.Device{
		ID:       primitive.NewObjectID(),
		UserID:   oid,
		DeviceID: d.DeviceId,
		Desc:     d.Desc,
		Time:     time.Now(),
	}

	log.Debug(device)

	doc, err := bson.Marshal(&device)
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

	c.JSON(util.Success(gin.H{"id": device.ID}))
}

func Devicelist(c *gin.Context) {
	limit, offset, err := util.Page(c)
	if err != nil {
		c.JSON(util.Error(400, "分页参数错误"))
		return
	}

	user := get_user(c)

	delete := false
	if c.Query("delete") == "true" {
		delete = true
	}

	collection := v1.GetDeviceColl()

	context := context.Background()

	count, err := collection.CountDocuments(context, bson.M{"user_id": user.ID, "delete": delete}, &options.CountOptions{})
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
		bson.M{"user_id": user.ID, "delete": delete},
		&options.FindOptions{Limit: &limit, Skip: &offset},
	)

	if err != nil {
		c.JSON(util.Error(500, err.Error()))
		return
	}

	defer cur.Close(context)

	arrays := make([]v1.Device, 0, 10)

	for cur.Next(context) {
		var device v1.Device
		if err := cur.Decode(&device); err != nil {
			c.JSON(util.Error(500, err.Error()))
			return
		}

		arrays = append(arrays, device)
	}

	c.JSON(util.Success(gin.H{
		"items": arrays,
		"count": count,
	}))
}

func DeviceDetail(c *gin.Context) {
	device_id := c.Param("id")

	if len(device_id) != 24 {
		c.JSON(util.Error(404, "404"))
		return
	}

	oid, err := primitive.ObjectIDFromHex(device_id)
	if err != nil {
		c.JSON(util.Error(500, err.Error()))
		return
	}

	user := get_user(c)

	delete := false
	if c.Query("delete") == "true" {
		delete = true
	}

	collection := v1.GetDeviceColl()

	result := collection.FindOne(context.Background(),
		bson.M{"_id": oid, "delete": delete},
		&options.FindOneOptions{})

	if result.Err() != nil {
		c.JSON(util.Error(500, result.Err().Error()))
		return
	}

	var device v1.Device
	if err = result.Decode(&device); err != nil {
		if err.Error() == "mongo: no documents in result" {
			c.JSON(util.Error(404, "404"))
			return
		}

		c.JSON(util.Error(500, err.Error()))
		return
	}

	if !user.Super && device.UserID != user.ID {
		c.JSON(util.Error(404, "404"))
		return
	}

	c.JSON(util.Success(device))
}

func DevicePost(c *gin.Context) {
	var d struct {
		DeviceId string `json:"device_id"`
		Desc     string `json:"desc"`
	}

	if err := c.Bind(&d); err != nil {
		c.JSON(util.Error(400, err.Error()))
		return
	}

	if d.DeviceId == "" {
		c.JSON(util.Error(400, "device_id 不能为空"))
		return
	}

	collection := v1.GetDeviceColl()

	count, err := collection.CountDocuments(context.Background(), bson.M{"device_id": d.DeviceId}, &options.CountOptions{})
	if err != nil {
		c.JSON(util.Error(500, err.Error()))
		return
	}

	if count > 0 {
		c.JSON(util.Error(400, "device_id 重复"))
		return
	}

	user := get_user(c)

	device := v1.Device{
		ID:       primitive.NewObjectID(),
		UserID:   user.ID,
		DeviceID: d.DeviceId,
		Desc:     d.Desc,
		Time:     time.Now(),
	}

	log.Debug(device)

	doc, err := bson.Marshal(&device)
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

	c.JSON(util.Success(gin.H{"id": device.ID}))
}

func DevicePatch(c *gin.Context) {
	device_id := c.Param("id")

	if len(device_id) != 24 {
		c.JSON(util.Error(404, "404"))
		return
	}

	oid, err := primitive.ObjectIDFromHex(device_id)
	if err != nil {
		c.JSON(util.Error(500, err.Error()))
		return
	}

	user := get_user(c)

	delete := false
	if c.Query("delete") == "true" {
		delete = true
	}

	collection := v1.GetDeviceColl()

	result := collection.FindOne(context.Background(), bson.M{"_id": oid, "delete": delete}, &options.FindOneOptions{})
	if result.Err() != nil {
		c.JSON(util.Error(500, result.Err().Error()))
		return
	}

	var device v1.Device
	if err = result.Decode(&device); err != nil {
		if err.Error() == "mongo: no documents in result" {
			c.JSON(util.Error(404, "404"))
			return
		}

		c.JSON(util.Error(500, err.Error()))
		return
	}

	if !user.Super && device.UserID != user.ID {
		c.JSON(util.Error(404, "404"))
		return
	}

	var d struct {
		DeviceId string `json:"device_id"`
		Desc     string `json:"desc"`
	}

	if err := c.Bind(&d); err != nil {
		c.JSON(util.Error(400, err.Error()))
		return
	}

	doc := bson.M{}

	if d.DeviceId != "" {
		count, err := collection.CountDocuments(context.Background(), bson.M{"device_id": d.DeviceId}, &options.CountOptions{})
		if err != nil {
			c.JSON(util.Error(500, err.Error()))
			return
		}

		if count > 0 && d.DeviceId != device.DeviceID {
			c.JSON(util.Error(400, "device_id 重复"))
			return
		}

		doc["device_id"] = d.DeviceId
	}

	if d.Desc != "" {
		doc["desc"] = d.Desc
	}

	if len(doc) > 0 {
		_, err = collection.UpdateOne(context.Background(), bson.M{"_id": oid}, bson.M{"$set": doc}, &options.UpdateOptions{})
		if err != nil {
			c.JSON(util.Error(500, err.Error()))
			return
		}
	}

	c.JSON(util.Success(gin.H{"id": device_id}))
}

func DeviceDelete(c *gin.Context) {
	device_id := c.Param("id")

	if len(device_id) != 24 {
		c.JSON(util.Error(404, "404"))
		return
	}

	oid, err := primitive.ObjectIDFromHex(device_id)
	if err != nil {
		c.JSON(util.Error(500, err.Error()))
		return
	}

	user := get_user(c)

	delete := false
	if c.Query("delete") == "true" {
		delete = true
	}

	collection := v1.GetDeviceColl()

	result := collection.FindOne(context.Background(), bson.M{"_id": oid, "delete": delete}, &options.FindOneOptions{})
	if result.Err() != nil {
		c.JSON(util.Error(500, result.Err().Error()))
		return
	}

	var device v1.Device
	if err = result.Decode(&device); err != nil {
		if err.Error() == "mongo: no documents in result" {
			c.JSON(util.Error(404, "404"))
			return
		}

		c.JSON(util.Error(500, err.Error()))
		return
	}

	if !user.Super && device.UserID != user.ID {
		c.JSON(util.Error(404, "404"))
		return
	}

	_, err = collection.UpdateOne(context.Background(), bson.M{"_id": oid}, bson.M{"$set": bson.M{"delete": true}}, &options.UpdateOptions{})
	if err != nil {
		c.JSON(util.Error(500, err.Error()))
		return
	}

	c.JSON(util.Success(gin.H{"id": device_id}))
}

func DeviceReset(c *gin.Context) {
	device_id := c.Param("id")

	if len(device_id) != 24 {
		c.JSON(util.Error(404, "404"))
		return
	}

	oid, err := primitive.ObjectIDFromHex(device_id)
	if err != nil {
		c.JSON(util.Error(500, err.Error()))
		return
	}

	user := get_user(c)

	delete := false
	if c.Query("delete") == "true" {
		delete = true
	}

	collection := v1.GetDeviceColl()

	result := collection.FindOne(context.Background(), bson.M{"_id": oid, "delete": delete}, &options.FindOneOptions{})
	if result.Err() != nil {
		c.JSON(util.Error(500, result.Err().Error()))
		return
	}

	var device v1.Device
	if err = result.Decode(&device); err != nil {
		if err.Error() == "mongo: no documents in result" {
			c.JSON(util.Error(404, "404"))
			return
		}

		c.JSON(util.Error(500, err.Error()))
		return
	}

	if !user.Super && device.UserID != user.ID {
		c.JSON(util.Error(404, "404"))
		return
	}

	_, err = collection.UpdateOne(context.Background(), bson.M{"_id": oid}, bson.M{"$set": bson.M{"delete": false}}, &options.UpdateOptions{})
	if err != nil {
		c.JSON(util.Error(500, err.Error()))
		return
	}

	c.JSON(util.Success(gin.H{"id": device_id}))
}
