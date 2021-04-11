// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	prices2 "github.com/jeb2239/diframeworks/lib/prices"
	"github.com/jeb2239/diframeworks/lib/processor"
)

// Injectors from wire.go:

func InitializeProcessor() *prices.Processor {
	iTimeProvider := makeTime()
	iStore := prices2.NewPricesStore(iTimeProvider)
	processor := prices.NewProcessorFromTimeProviderAndPriceStore(iTimeProvider, iStore)
	return processor
}