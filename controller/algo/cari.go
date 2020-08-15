package algo

import (
	"math/rand"
	"strings"
)

func (a Algo) Random(max int64, min int64) int64 { return rand.Int63n(max-min) + min }

func (a Algo) AccurateSearch(title string, payload string, ulasan uint32, terjual uint32, rating uint16, key []string) float64 {

	litTitle := strings.ToLower(title)
	litPayload := strings.ToLower(payload)

	//fmt.Println(litTitle)
	//fmt.Println(litPayload)

	tmpArr := strings.Split(litPayload, "%20")
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
		sta := strings.Contains(litTitle, tmpArr[i])
		//fmt.Println(sta)
		if sta == true {
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
