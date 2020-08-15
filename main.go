package main

import (
	"brg/router"
	"fmt"
	"runtime"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
)

const port = 3000

var r router.Router

func init() {
	r = router.Router{}
}

func main() {

	runtime.GOMAXPROCS(4)

	app := fiber.New()

	app.Use(cors.New())

	fmt.Println("Service Barang [RUN] ", port)

	app.Get("v1/api/barang/page", r.GetGoceng)
	app.Get("v1/api/barang/:idBarang", r.GetBarang)
	app.Get("v1/api/barang/cari/:cari", r.CariBarang)

	app.Post("v1/api/barang/insert", r.InsertBarang)

	app.Put("v1/api/barang/:idBarang", r.UpdateBarang)
	app.Put("v1/api/barang/stok/:idBarang", r.UpdateStokBarang)
	app.Put("v1/api/barang/nama/:idBarang", r.UpdateNamaBarang)
	app.Put("v1/api/barang/deskripsi/:idBarang", r.UpdateDeskripsiBarang)
	app.Put("v1/api/barang/total_ulasan/:idBarang", r.UpdateTotalUlasanBarang)
	app.Put("v1/api/barang/total_terjual/:idBarang", r.UpdateTotalTerjualBarang)
	app.Put("v1/api/barang/price/:idBarang", r.UpdatePriceBarang)
	app.Put("v1/api/barang/keyword/:idBarang", r.UpdateKeywordBarang)
	app.Put("v1/api/barang/rating/:idBarang", r.UpdateRatingBarang)

	app.Listen(port)

}
