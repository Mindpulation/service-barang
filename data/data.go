package data

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ParamIdBarang struct {
	Id primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
}

type DataUlasan struct {
	Nama   string
	Ulasan string
	Rating float32
}

type DataStok struct {
	Stok uint64
}

type DataBarang struct {
	Id           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IdToko       string
	KodeBarcode  string
	Nama         string
	Deskripsi    string
	TotalUlasan  uint64
	TotalTerjual uint64
	Keyword      []string
	Rating       uint64
	Price        uint64
	Stok         uint32
	Ulasan       []DataUlasan
}

type DataFilter struct {
	Id           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IdToko       string
	KodeBarcode  string
	Nama         string
	Deskripsi    string
	TotalUlasan  uint64
	TotalTerjual uint64
	Keyword      []string
	Rating       uint64
	Price        uint64
	Stok         uint32
	Ulasan       []DataUlasan
	FilterCari   int
}
