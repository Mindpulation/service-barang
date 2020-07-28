package algo

import (
	"brg/controller"
	"brg/model"
)

type Algo struct{}

var (
	batasAkhir, banyakIndex, limitIndex int64
	max                                 int64 = 5000
	m                                   model.MongoDB
	c                                   controller.Controller
)

func init() {
	m = model.MongoDB{"mongodb://127.0.0.1:27017", "BarangDB", "Barang"}
	c = controller.Controller{}
}
