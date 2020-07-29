package model

import "go.mongodb.org/mongo-driver/mongo"

func (m MongoDB) InsertOne(col *mongo.Collection, e interface{}) error {
	ct := m.MakeContext(5)
	_, err := col.InsertOne(ct, e)
	return err
}
