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

type DataNama struct {
	Nama string
}

type DataDeskripsi struct {
	Deskripsi string
}

type DataTotalUlasan struct {
	TotalUlasan uint64
}

type DataTotalTerjual struct {
	TotalTerjual uint64
}

type DataPrice struct {
	Price uint64
}

type DataKeyword struct {
	Keyword []string
}

type DataRating struct {
	Rating float64
}

type DataBarang struct {
	Id           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IdToko       string             `json:"idToko,omitempty" bson:"idToko,omitempty"`
	KodeBarcode  string             `json:"barcode,omitempty" bson:"barcode,omitempty"`
	Nama         string             `json:"name,omitempty" bson:"name,omitempty"`
	Deskripsi    string             `json:"deskripsi,omitempty" bson:"deskripsi,omitempty"`
	TotalUlasan  uint32             `json:"totalUlasan,omitempty" bson:"totalUlasan,omitempty"`
	TotalTerjual uint32             `json:"totalTerjual,omitempty" bson:"totalTerjual,omitempty"`
	Keyword      []string           `json:"key,omitempty" bson:"key,omitempty"`
	Rating       uint16             `json:"rating,omitempty" bson:"rating,omitempty"`
	Price        uint64             `json:"price,omitempty" bson:"price, omitempty"`
	Stok         uint32             `json:"stok,omitempty" bson:"stok, omitempty"`
	Ulasan       []DataUlasan       `json:"ulasan" bson:"ulasan"`
}

type DataFilter struct {
	Id           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	IdToko       string             `json:"idToko,omitempty" bson:"idToko,omitempty"`
	KodeBarcode  string             `json:"barcode,omitempty" bson:"barcode,omitempty"`
	Nama         string             `json:"name,omitempty" bson:"name,omitempty"`
	Deskripsi    string             `json:"deskripsi,omitempty" bson:"deskripsi,omitempty"`
	TotalUlasan  uint32             `json:"totalUlasan,omitempty" bson:"totalUlasan,omitempty"`
	TotalTerjual uint32             `json:"totalTerjual,omitempty" bson:"totalTerjual,omitempty"`
	Keyword      []string           `json:"key,omitempty" bson:"key,omitempty"`
	Rating       uint16             `json:"rating,omitempty" bson:"rating,omitempty"`
	Price        uint64             `json:"price,omitempty" bson:"price, omitempty"`
	Stok         uint32             `json:"stok,omitempty" bson:"stok, omitempty"`
	Ulasan       []DataUlasan       `json:"ulasan" bson:"ulasan"`
	FilterCari   float64            `json:"filter" bson:"fitler"`
}
