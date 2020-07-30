package algo

import (
	"brg/data"
	"math/rand"
	"strings"
	"sync"
)

func (a Algo) Random(max int64, min int64) int64 { return rand.Int63n(max-min) + min }

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

func (a Algo) GetAccurateData() []data.DataBarang {
	skip, lim := a.SkipLimitProduct()

	cn, cl, _ := m.MakeConnection()
	cr, _ := m.FindSkipAndLimit(cl, skip, lim)
	arr := c.ToArray(cr)
	m.Disconnect(cn)

	return arr
}

func (a Algo) AccurateSearch(title string, payload string, ulasan uint32, terjual uint32, rating uint16, key []string) float64 {

	litTitle := strings.ToLower(title)
	litPayload := strings.ToLower(payload)

	tmpArr := strings.Split(litPayload, " ")
	leng := len(tmpArr) // Banyak kata bedasarkan kata pencarian
	lengKey := len(key) // Banyak kata pada sebuah field Keyword
	i := 0
	j := 0

	var (
		nilaiJudul, persenRating, persenTerjual, persenUlasan, persenJudul, persenKey, total float64
		nilaiKeyword                                                                         bool
	)

	persenRating = (float64(rating) / 5) * 20

	if terjual >= 100 {
		persenTerjual = float64(8)
		persenUlasan = float64(7)
	} else {
		persenTerjual = (float64(terjual) / 100) * 8
		persenUlasan = (float64(ulasan) / 100) * 7
	}

	for i < leng {
		if strings.Contains(litTitle, tmpArr[i]) {
			nilaiJudul++
		}
		i++
	}

	persenJudul = (nilaiJudul / float64(leng)) * 50

	for j < lengKey {
		if strings.Contains(litPayload, strings.ToLower(key[j])) {
			nilaiKeyword = true
			break
		}
		j++
	}

	if persenJudul == 0 {

		if nilaiKeyword == true {
			persenKey = 50
		}

	} else {

		if nilaiKeyword == true {
			persenKey = 15
		}

	}

	total = persenJudul + persenRating + persenTerjual + persenUlasan + persenKey
	return total
}

func (a Algo) MergerFilter(e data.DataBarang, filter float64) data.DataFilter {
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

func (a Algo) CalculateIndex(leng int) (int, int, int) {
	max := 500
	countIndex := (leng / max) - 1
	countLimitIndex := countIndex + 1
	endData := leng % max

	return countIndex, countLimitIndex, endData
}

func (a Algo) RangeIndex(index int, limit int, endData int) (int, int) {
	var (
		start, end int
	)
	start = index * 500
	if index == limit {
		end = start + endData
	} else {
		end = start + 500
	}
	return start, end
}

func (a Algo) PartOfSearch(payload string, arr []data.DataBarang, aim *[]data.DataFilter, wg *sync.WaitGroup) {

	var objFilter data.DataFilter
	var arrFilter []data.DataFilter

	for _, e := range arr {
		nilai := a.AccurateSearch(e.Nama, payload, e.TotalUlasan, e.TotalTerjual, e.Rating, e.Keyword)
		if nilai >= 50 {
			objFilter = a.MergerFilter(e, nilai)
			arrFilter = append(arrFilter, objFilter)
		}
	}

	*aim = append(*aim, arrFilter...)

	wg.Done()

}

func (a Algo) LoopSearch(payload string) []data.DataFilter {

	var (
		arrFilter  []data.DataFilter
		objFilter  data.DataFilter
		wg         sync.WaitGroup
		start, end int
	)

	arr := a.GetAccurateData()

	leng := len(arr)

	if leng < 500 {

		for _, e := range arr {
			nilai := a.AccurateSearch(e.Nama, payload, e.TotalUlasan, e.TotalTerjual, e.Rating, e.Keyword)
			if nilai >= 50 {
				objFilter = a.MergerFilter(e, nilai)
				arrFilter = append(arrFilter, objFilter)
			}
		}

		return arrFilter

	}

	i := 0

	lengIndex, limIndex, endData := a.CalculateIndex(leng)

	for i <= lengIndex+1 {
		start, end = a.RangeIndex(i, limIndex, endData)
		wg.Add(1)
		go a.PartOfSearch(payload, arr[start:end], &arrFilter, &wg)
		wg.Wait()
		i++
	}

	return arrFilter

}
