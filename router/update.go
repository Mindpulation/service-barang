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

func (r Router) UpdateNamaBarang(c *fiber.Ctx) {
	con, col, _ := m.MakeConnection()

	var nama data.DataNama
	json.Unmarshal([]byte(c.Body()), &nama)

	update := bson.M{
		"$set": nama,
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

func (r Router) UpdateDeskripsiBarang(c *fiber.Ctx) {
	con, col, _ := m.MakeConnection()

	var deskripsi data.DataDeskripsi
	json.Unmarshal([]byte(c.Body()), &deskripsi)

	update := bson.M{
		"$set": deskripsi,
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

func (r Router) UpdateTotalUlasanBarang(c *fiber.Ctx) {
	con, col, _ := m.MakeConnection()

	var totalUlasan data.DataTotalUlasan
	json.Unmarshal([]byte(c.Body()), &totalUlasan)

	update := bson.M{
		"$set": totalUlasan,
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

func (r Router) UpdateTotalTerjualBarang(c *fiber.Ctx) {
	con, col, _ := m.MakeConnection()

	var totalTerjual data.DataTotalTerjual
	json.Unmarshal([]byte(c.Body()), &totalTerjual)

	update := bson.M{
		"$set": totalTerjual,
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

func (r Router) UpdatePriceBarang(c *fiber.Ctx) {
	con, col, _ := m.MakeConnection()

	var price data.DataStok
	json.Unmarshal([]byte(c.Body()), &price)

	update := bson.M{
		"$set": price,
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

func (r Router) UpdateKeywordBarang(c *fiber.Ctx) {
	con, col, _ := m.MakeConnection()

	var keyword data.DataKeyword
	json.Unmarshal([]byte(c.Body()), &keyword)

	update := bson.M{
		"$set": keyword,
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

func (r Router) UpdateRatingBarang(c *fiber.Ctx) {
	con, col, _ := m.MakeConnection()

	var rating data.DataRating
	json.Unmarshal([]byte(c.Body()), &rating)

	update := bson.M{
		"$set": rating,
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
