//+build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"github.com/google/wire"

	"github.com/jeb2239/diframeworks/lib/prices"
	proc "github.com/jeb2239/diframeworks/lib/processor"
)

// InitializeEvent creates an Event. It will error if the Event is staffed with
// a grumpy greeter.
func InitializeProcessor() *proc.Processor {

	wire.Build(prices.NewPricesStore,
		proc.NewProcessorFromTimeProviderAndPriceStore,
		makeTime)
	return &proc.Processor{}
}
