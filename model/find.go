package model

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m MongoDB) FindAll(col *mongo.Collection) (*mongo.Cursor, error) {
	ct := m.MakeContext(5)
	res, err := col.Find(ct, bson.M{})
	return res, err
}

func (m MongoDB) FindSkipAndLimit(col *mongo.Collection, skip int64, limit int64) (*mongo.Cursor, error) {
	ct := m.MakeContext(100)
	option := options.FindOptions{}
	return col.Find(ct, bson.M{}, option.SetSkip(skip), option.SetLimit(limit))
}

func (m MongoDB) FindCount(col *mongo.Collection) (int64, error) {
	ct := m.MakeContext(5)
	res, err := col.CountDocuments(ct, bson.M{})
	fmt.Println(err)
	return res, err
}

func (m MongoDB) FindWithParam(col *mongo.Collection, filter interface{}) (*mongo.Cursor, error) {
	ct := m.MakeContext(5)
	res, err := col.Find(ct, filter)
	return res, err
}
