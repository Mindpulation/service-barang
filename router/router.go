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
<<<<<<< HEAD
	m = model.MongoDB{"mongodb://mongo:27017/", "BarangDB", "Barang"}
=======
	m = model.MongoDB{"mongodb://mongo-indo1:27017/BarangDB", "BarangDB", "Barang"}
>>>>>>> 6e10aea4fcf84399c325a4d0aca406039d80ae57
	cl = controller.Controller{}
}
