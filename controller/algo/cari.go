package algo

import (
	"brg/data"
	"math/rand"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

func (a Algo) Random(max int64, min int64) int64 {
	return rand.Int63n(max-min) + min
}

func (a Algo) TagCalculate(qty int64) int64 {
	batasAkhir = qty % max
	banyakIndex = qty / max
	limitIndex = banyakIndex + 1
	return a.Random(limitIndex, 0) + 1
}

func (a Algo) CalculateRange(index int64) (int64, int64) {

	var (
		start int64
		end   int64
	)

	start = (index - 1) * max

	if index <= banyakIndex {
		end = (index * max) - 1
		return start, end
	} else if index == limitIndex {
		end = start + batasAkhir
		return start, end
	} else {
		return 0, 0
	}

}

func (a Algo) SkipLimitProduct() (int64, int64) {
	cn, cl, _ := m.MakeConnection()
	qty, _ := m.FindCount(cl)

	var (
		skip, limit, akhir int64
	)

	if qty > max {
		idx := a.TagCalculate(qty)
		skip, akhir = a.CalculateRange(idx)
		limit = akhir - skip
	} else {
		skip = 0
		akhir = qty
		limit = akhir
	}

	m.Disconnect(cn)

	return skip, limit
}

func (a Algo) GetAccurateData() *mongo.Cursor {
	skip, lim := a.SkipLimitProduct()
	cn, cl, _ := m.MakeConnection()
	cr, _ := m.FindSkipAndLimit(cl, skip, lim)
	m.Disconnect(cn)

	return cr
}

func (a Algo) AccurateSearch(title string, payload string, ulasan uint32, terjual uint32, rating uint16, key []string) int {

	tmpArr := strings.Split(payload, " ")

	leng := len(tmpArr)
	i := 0

	lengKey := len(key)
	j := 0

	// Judul   => 50%
	// Rating  => 20%
	// Keyword => 15%
	// Terjual => 8%
	// Ulasan  => 7%

	nilaiJudul := 0
	nilaiKeyword := 0

	persenRating := int((rating / 5) * 20)

	var persenTerjual int

	if terjual >= 100 {
		persenTerjual = int(8)
	} else {
		persenTerjual = int((terjual / 100) * 8)
	}

	var persenUlasan int

	if terjual >= 100 {
		persenUlasan = int(7)
	} else {
		persenUlasan = int((ulasan / 100) * 7)
	}

	for i < leng {
		if strings.Contains(title, tmpArr[i]) {
			nilaiJudul++
		}
	}

	persenJudul := int((nilaiJudul / leng) * 50)

	for j < lengKey {
		if strings.Contains(payload, key[j]) {
			nilaiKeyword++
		}
	}

	persenKey := int((nilaiKeyword / lengKey) * 15)

	total := persenJudul + persenRating + persenTerjual + persenUlasan + persenKey

	return total

}

func (a Algo) MergerFilter(e data.DataBarang, filter int) data.DataFilter {
	var res data.DataFilter
	res.Id = e.Id
	res.IdToko = e.IdToko
	res.KodeBarcode = e.KodeBarcode
	res.Nama = e.Nama
	res.Deskripsi = e.Deskripsi
	res.TotalUlasan = e.TotalUlasan
	res.TotalTerjual = e.TotalTerjual
	res.Keyword = e.Keyword
	res.Rating = e.Rating
	res.Price = e.Price
	res.Stok = e.Stok
	res.Ulasan = e.Ulasan
	res.FilterCari = filter
	return res
}

func (a Algo) LoopSearch(title string) []data.DataFilter {
	cr := a.GetAccurateData()
	arr := c.ToArray(cr)

	var arrFilter []data.DataFilter
	var objFilter data.DataFilter

	for _, e := range arr {
		nilai := a.AccurateSearch(title, e.Nama, e.TotalUlasan, e.TotalTerjual, e.Rating, e.Keyword)
		if nilai >= 50 {
			objFilter = a.MergerFilter(e, nilai)
			arrFilter = append(arrFilter, objFilter)
		}
	}

	return arrFilter
}
