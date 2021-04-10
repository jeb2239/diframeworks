package prices

import (
	"time"

	"github.com/jeb2239/diframeworks/lib/chrono"
)

type PriceRecord struct {
	Price     int
	Name      string
	TimeStamp time.Time
}


type IStore interface {
	GetPrices() []PriceRecord
}

type Store struct {
	timeProvider chrono.ITimeProvider
}

func NewPricesStore(tp chrono.ITimeProvider) IStore {

	return &Store{
		timeProvider: tp,
	}

}

func (s *Store) GetPrices() []PriceRecord {

	return []PriceRecord{
		{
			Price:     22,
			Name:      "IBM",
			TimeStamp: s.timeProvider.GetTime(),
		},
	}
}
