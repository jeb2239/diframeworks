package main

import (
	"fmt"

	"github.com/jeb2239/diframeworks/lib/chrono"
	"github.com/jeb2239/diframeworks/lib/prices"
	proc "github.com/jeb2239/diframeworks/lib/processor"
	"go.uber.org/fx"
)

func main() {

	makeTime := func() chrono.ITimeProvider {
		return &chrono.TimeProvider{}
	}

	app := fx.New(
		fx.Provide(
			prices.NewPricesStore,
			proc.NewProcessorFromTimeProviderAndPriceStore,
			makeTime,
		),
		fx.Invoke(
			func(stuffDoer *proc.Processor) {
				fmt.Println("==stuffDoer is being run==")
				stuffDoer.DoIO()
			},
		),
	)
	app.Run()

}
