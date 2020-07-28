package controller

import (
	"brg/data"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Controller struct {
}

func (c Controller) ToArray(cr *mongo.Cursor) []data.DataBarang {
	ct, _ := context.WithTimeout(context.Background(), 100*time.Second)
	var arrReturn []data.DataBarang
	var objectBarang data.DataBarang
	for cr.Next(ct) {
		cr.Decode(&objectBarang)
		arrReturn = append(arrReturn, objectBarang)
	}

	return arrReturn
}
