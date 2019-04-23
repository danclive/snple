package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID     primitive.ObjectID `bson:"_id" json:"id"`
	Name   string             `bson:"name" json:"name"`
	Pass   string             `bson:"pass" json:"-"`
	Time   primitive.DateTime `bson:"time" json:"time"`
	Super  bool               `bson:"super" json:"super"`
	Delete bool               `bson:"delete" json:"-"`
}

type Device struct {
	ID       primitive.ObjectID `bson:"_id"`
	UserID   primitive.ObjectID `bson:"_id"`
	DeviceID string             `bson:"device_id"`
	Desc     string             `bson:"desc"`
	Status   int32              `bson:"status"`
}
