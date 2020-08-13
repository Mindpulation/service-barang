package router

import (
	"brg/controller"
	"brg/controller/algo"
	"brg/model"
)

type Router struct {
}

type ResBool struct {
	sta bool
	err error
}

var a algo.Algo

var bol ResBool

var m model.MongoDB
var cl controller.Controller

func init() {
	m = model.MongoDB{"mongodb://mongo-indo1:27017/BarangDB", "BarangDB", "Barang"}
	cl = controller.Controller{}
}
