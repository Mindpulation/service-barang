package algo

import (
	"brg/controller"
	"brg/model"
)

type Algo struct{}

var m model.MongoDB

var c controller.Controller

var max int64 = 5000

var batasAkhir int64

var banyakIndex int64

var limitIndex int64

func init() {
	m = model.MongoDB{"mongodb://127.0.0.1:27017", "BarangDB", "Barang"}
	c = controller.Controller{}
}
