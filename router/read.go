package router

import (
	"brg/data"
	"encoding/json"

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

	cari := c.Params("cari")

	res := a.LoopSearch(cari)

	// var filterRes []data.DataFilter

	// for i, e := range res {
	// 	filterRes = append(filterRes, e)
	// 	if i == 99 {
	// 		break
	// 	}
	// }

	response, _ := json.Marshal(res)
	c.Send(response)
}
