package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m MongoDB) UpdOne(col *mongo.Collection, find interface{}, set primitive.D) error {
	ct := m.MakeContext(5)
	_, err := col.UpdateOne(ct, find, set)
	return err
}
