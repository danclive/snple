package v1

import (
	"time"

	"github.com/danclive/mqtt-console/config"
	"github.com/danclive/mqtt-console/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID     primitive.ObjectID `bson:"_id" json:"id"`
	Name   string             `bson:"name" json:"name"`
	Pass   string             `bson:"pass" json:"-"`
	Time   time.Time          `bson:"time" json:"time"`
	Super  bool               `bson:"super" json:"super"`
	Delete bool               `bson:"delete" json:"-"`
	Desc   string             `bson:"desc" json:"desc"`
}

func (*User) ColName() string {
	return "user"
}

func GetUserColl() *mongo.Collection {
	return db.MongoClient.Database(config.Config.Mongo.Database, nil).Collection("user", nil)
}

type Device struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	UserID   primitive.ObjectID `bson:"user_id" json:"user_id"`
	DeviceID string             `bson:"device_id" json:"device_id"`
	Desc     string             `bson:"desc" json:"desc"`
	Status   int32              `bson:"status" json:"status"` // 0: ready, 1: active, 2: error, 3: temp
	Delete   bool               `bson:"delete" json:"-"`
	Time     time.Time          `bson:"time" json:"time"`
	// store bool `bson:"store" json:"store"` // 是否保存数据
	// values Values `bson:"values" json:"values"`
}

type Values struct {
	Name string    `bson:"name"`
	Desc string    `bson:"desc"`
	Type string    `bson:"type"` // int, float, bool, string
	Time time.Time `bson:"time" json:"time"`
}

func (*Device) ColName() string {
	return "device"
}

func GetDeviceColl() *mongo.Collection {
	return db.MongoClient.Database(config.Config.Mongo.Database, nil).Collection("device", nil)
}

type Option struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	Lable string             `bson:"lable" json:"lable"`
	Name  string             `bson:"name"`
	Value string             `bson:"value"`
	VType string             `bson:"vtype"`
}

func (*Option) ColName() string {
	return "option"
}

func GetOptionColl() *mongo.Collection {
	return db.MongoClient.Database(config.Config.Mongo.Database, nil).Collection("option", nil)
}
