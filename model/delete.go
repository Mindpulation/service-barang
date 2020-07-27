package model

import "go.mongodb.org/mongo-driver/mongo"

func (m MongoDB) DelOne(col *mongo.Collection, e interface{}) error {
	ct := m.MakeContext(5)
	_, err := col.DeleteOne(ct, e)
	return err
}
