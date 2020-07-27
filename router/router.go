package router

import (
	"brg/controller"
	"brg/model"
)

type Router struct {
}

type ResBool struct {
	sta bool
	err error
}

var bol ResBool

var m model.MongoDB
var cl controller.Controller

func init() {
	m = model.MongoDB{"mongodb://localhost:27017", "BarangDB", "Barang"}
	cl = controller.Controller{}
}
