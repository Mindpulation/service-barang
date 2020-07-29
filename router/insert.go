package router

import (
	"brg/data"
	"encoding/json"
	"github.com/gofiber/fiber"
)

func (r Router) InsertBarang(c *fiber.Ctx) {
	con, col, _ := m.MakeConnection()

	var barang data.DataBarang
	json.Unmarshal([]byte(c.Body()), &barang)

	err := m.InsertOne(col, barang)

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
