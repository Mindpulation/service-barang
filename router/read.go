package router

import (
	"brg/data"
	"encoding/json"
	"sync"

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

	res := a.CariBarang(cari)

	var response []byte

	if res == nil {
		response, _ = json.Marshal(res)
	} else {
		response, _ = json.Marshal(res[0:25])
	}

	c.Send(response)
}

func (r Router) GetGoceng(c *fiber.Ctx) {

	//start := time.Now().Nanosecond()

	con, col, _ := m.MakeConnection()

	var res []data.DataBarang

	var wg sync.WaitGroup

	var getRange = func(index int64) int64 {
		return (index * 500)
	}

	var getData = func(index int64, wg *sync.WaitGroup, res *[]data.DataBarang) {
		s := getRange(index)
		cr, _ := m.FindSkipAndLimit(col, s, 500)
		t := cl.ToArray(cr)
		*res = append(*res, t...)
		defer wg.Done()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go getData(int64(i), &wg, &res)
	}

	wg.Wait()

	// cr, _ := m.FindSkipAndLimit(col, 0, 5000)
	// res = cl.ToArray(cr)

	m.Disconnect(con)

	//end := time.Now().Nanosecond() - start

	//fmt.Println(float64(end) / 1000000)

	response, _ := json.Marshal(res)
	c.Send(response)
}
