package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/danclive/mqtt-console/config"
	"github.com/danclive/mqtt-console/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func InitMongo() {
	log.Info("连接 mongo ...")
	client, err := mongo.NewClient(options.Client().ApplyURI(config.Config.Mongo.Uri))
	if err != nil {
		log.Fatal("连接 mongo 失败: ", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("连接 mongo 失败: ", err)
	}

	log.Info("OK！")

	log.Info("ping mongo ...")
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("ping mongo 失败: ", err)
	}

	MongoClient = client

	log.Info("OK！")

	initDatabse()
}

func initDatabse() {
	log.Info("初始化数据库 ...")

	collection := MongoClient.Database(config.Config.Mongo.Database, nil).Collection("user", nil)

	count, err := collection.CountDocuments(context.Background(), bson.D{}, nil)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		user := User{
			ID:    primitive.NewObjectID(),
			Name:  "admin",
			Pass:  "efa1f375d76194fa51a3556a97e641e61685f914d446979da50a551a4333ffd7",
			Time:  time.Now(),
			Super: true,
			Desc:  "管理员",
		}

		doc, err := bson.Marshal(&user)
		if err != nil {
			log.Fatal(err)
		}

		log.Debug(doc)

		result, err := collection.InsertOne(context.Background(), doc, &options.InsertOneOptions{})
		if err != nil {
			log.Fatal(err)
		}

		log.Debug(result)
	}

	log.Info("OK！")
}
