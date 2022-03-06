package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MongoConfig struct {
	*options.ClientOptions
	Database string
}

type MongoHelper struct {
	client *mongo.Client
	conf   *MongoConfig
	db     *mongo.Database
}

func NewMongoHelper(conf *MongoConfig) *MongoHelper {
	return &MongoHelper{conf: conf}
}

func (helper *MongoHelper) Collection(name string) *mongo.Collection {
	return helper.client.Database(helper.conf.Database).Collection(name)
}

func (helper *MongoHelper) OpenConnection() error {
	if helper.client != nil {
		return nil
	}
	var err error
	options := helper.conf.ClientOptions
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	helper.client, err = mongo.Connect(ctx, options)
	if err != nil {
		log.Printf("%v\n", err)
		return err
	}
	err = helper.client.Ping(context.TODO(), nil)
	if err != nil {
		defer helper.closeConnection()
		log.Printf("%v\n", err)
		return err
	}
	log.Println("Connected to Mongo!")
	return nil
}

func (helper *MongoHelper) closeConnection() error {
	err := helper.client.Disconnect(context.TODO())
	if err != nil {
		return err
	}
	log.Println("Connected to Mongo closed.")
	return nil
}
