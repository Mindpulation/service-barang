package algo

import (
	"brg/data"
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func (a Algo) CariBarang(e string) []data.DataFilter {

	var (
		dat    []data.DataFilter
		wait   sync.WaitGroup
		en     int64
		qli    int64
		qRate  int64
		qMax   int64 = 5000
		qCount int64 = 500
	)

	cn, cl, _ := m.MakeConnection()
	q, _ := m.FindCount(cl)

	en = q % qMax

	qli = q / qMax

	qRate = qMax / qCount

	ran := a.Random(qli, 0)

	ss, ee := GetRange(ran, qli, en, qMax)

	var ToArray = func(cr *mongo.Cursor) {
		ct, _ := context.WithTimeout(context.Background(), 10*time.Second)

		var obj data.DataFilter

		for cr.Next(ct) {
			cr.Decode(&obj)

			val := a.AccurateSearch(obj.Nama, e, obj.TotalUlasan, obj.TotalTerjual, obj.Rating, obj.Keyword)

			if val > 50 {
				obj.FilterCari = float64(val)
				dat = append(dat, obj)
			}

		}

	}

	if ee < 500 {
		return nil
		// Exec
	}

	var Filtering = func(index int64) {
		s := Range(index, ss, qCount)
		cr, _ := m.FindSkipAndLimit(cl, s, qCount)
		ToArray(cr)
		defer wait.Done()
	}

	for i := 0; i < int(qRate); i++ {
		wait.Add(1)
		go Filtering(int64(i))
	}

	wait.Wait()

	m.Disconnect(cn)

	return dat
}

func GetRange(i int64, qli int64, en int64, max int64) (int64, int64) {
	s := i * max
	if qli == i {
		return s, en
	}
	return s, max
}

func Range(index int64, ss int64, qCount int64) int64 {
	return (index * qCount) + ss
}
