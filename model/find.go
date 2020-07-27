package model

import (

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m MongoDB) FindAll(col *mongo.Collection) (*mongo.Cursor, error){
	ct := m.MakeContext(5)	
	res, err := col.Find(ct, bson.M{})
	return res, err
}

func (m MongoDB) FindWithParam(col *mongo.Collection ,filter interface{}) (*mongo.Cursor, error){
	ct := m.MakeContext(5)			
	res, err := col.Find(ct, filter);
	return res, err
}