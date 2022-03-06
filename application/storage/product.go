package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"walmart-api/application/entities"
	"walmart-api/helpers/database"
)

type ProductStorage struct {
	mongo *database.MongoHelper
}

func NewProductStorage(mongo *database.MongoHelper) *ProductStorage {
	return &ProductStorage{mongo: mongo}
}

func (storage *ProductStorage) Find(id int64) ([]entities.Product, error) {
	var products []entities.Product
	err := storage.mongo.OpenConnection()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	filter := bson.D{bson.E{
		Key:   "id",
		Value: id,
	}}
	opts := &options.FindOptions{}
	cursor, err := storage.mongo.Collection(collectionProducts).Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var product entities.Product
		err = cursor.Decode(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
