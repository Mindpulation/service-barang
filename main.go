package main

import (
	"brg/router"
	"runtime"

	"github.com/gofiber/fiber"
)

const port = 2020

var r router.Router

func init() {
	r = router.Router{}
}

func main() {

	runtime.GOMAXPROCS(4)

	app := fiber.New()
	app.Get("v1/api/barang/:idBarang", r.GetBarang)
	app.Get("v1/api/barang/cari/:cari", r.CariBarang)
	app.Post("v1/api/barang/insert", r.InsertBarang)
	app.Put("v1/api/barang/:idBarang", r.UpdateBarang)
	app.Put("v1/api/barang/stok/:idBarang", r.UpdateStokBarang)

	app.Listen(port)
}
