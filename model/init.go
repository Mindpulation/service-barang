package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Server   string
	Database string
	Col      string
}

func (m MongoDB) MakeContext(e int) context.Context {
	duration := time.Duration(e) * time.Second
	ctx, _ := context.WithTimeout(context.Background(), duration)
	return ctx
}

func (m MongoDB) MakeConnection() (*mongo.Client, *mongo.Collection, error) {
	ct := m.MakeContext(5)
	con, err := mongo.Connect(ct, options.Client().ApplyURI(m.Server))
	if err != nil {
		return nil, nil, err
	}
	col := con.Database(m.Database).Collection(m.Col)
	return con, col, nil
}

func (m MongoDB) Disconnect(con *mongo.Client) {
	ct := m.MakeContext(5)
	con.Disconnect(ct)
}
