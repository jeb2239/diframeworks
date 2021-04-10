package prices

import (
	"fmt"

	"github.com/jeb2239/diframeworks/lib/chrono"
	"github.com/jeb2239/diframeworks/lib/prices"
)

type Processor struct {
	timeProvider chrono.ITimeProvider
	priceStore   prices.IStore
}

func NewProcessorFromTimeProviderAndPriceStore(tp chrono.ITimeProvider, ps prices.IStore) *Processor {
	fmt.Println("helo")
	return &Processor{
		timeProvider: tp,
		priceStore:   ps,
	}

}

func (p *Processor) DoIO() {
	values := p.priceStore.GetPrices()
	for _, v := range values {
		fmt.Println(v)
	}
}
