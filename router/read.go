package router

import (
	"brg/data"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r Router) GetBarang(c *fiber.Ctx) {
	con, col, _ := m.MakeConnection()
	paramsId := c.Params("idBarang")

	var objId, _ = primitive.ObjectIDFromHex(paramsId)
	var e data.ParamIdBarang

	e.Id = objId
	cr, _ := m.FindWithParam(col, e)

	res := cl.ToArray(cr)
	m.Disconnect(con)

	response, _ := json.Marshal(res)
	c.Send(response)
}

func (r Router) CariBarang(c *fiber.Ctx) {

	start := time.Now().Nanosecond()

	cari := c.Params("cari")

	res := a.LoopSearch(cari)

	end := time.Now().Nanosecond() - start

	fmt.Println(float64(end) / 1000000)

	response, _ := json.Marshal(res[0:25])
	c.Send(response)
}
