package controller

import (
	"brg/data"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Controller struct {
}

func (c Controller) ToArray(cr *mongo.Cursor) []data.DataBarang {
	ct, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var arrReturn []data.DataBarang
	var objectBarang data.DataBarang
	for cr.Next(ct) {
		cr.Decode(&objectBarang)
		arrReturn = append(arrReturn, objectBarang)
	}

	return arrReturn
}
