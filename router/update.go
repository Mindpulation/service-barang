package router

import (
	"brg/data"
	"encoding/json"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r Router) UpdateBarang(c *fiber.Ctx) {
	con, col, _ := m.MakeConnection()

	var barang data.DataBarang
	json.Unmarshal([]byte(c.Body()), &barang)

	update := bson.M{
		"$set": barang,
	}

	_id, _ := primitive.ObjectIDFromHex(c.Params("idBarang"))
	var Id data.ParamIdBarang
	Id.Id = _id

	err := m.UpdOne(col, Id, update)

	if err != nil {
		bol.sta = false
		bol.err = err
		c.Status(500).Send(bol)
		return
	} else {
		bol.sta = true
		bol.err = nil
	}
	m.Disconnect(con)
	response, _ := json.Marshal(bol)
	c.Send(response)
}

func (r Router) UpdateStokBarang(c *fiber.Ctx) {
	con, col, _ := m.MakeConnection()

	var stok data.DataStok
	json.Unmarshal([]byte(c.Body()), &stok)

	update := bson.M{
		"$set": stok,
	}

	_id, _ := primitive.ObjectIDFromHex(c.Params("idBarang"))
	var Id data.ParamIdBarang
	Id.Id = _id

	err := m.UpdOne(col, Id, update)

	if err != nil {
		bol.sta = false
		bol.err = err
		c.Status(500).Send(bol)
		return
	} else {
		bol.sta = true
		bol.err = nil
	}

	m.Disconnect(con)

	response, _ := json.Marshal(bol)
	c.Send(response)
}
