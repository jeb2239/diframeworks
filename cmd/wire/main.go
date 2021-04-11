package main

import "github.com/jeb2239/diframeworks/lib/chrono"

func makeTime() chrono.ITimeProvider {
	return &chrono.TimeProvider{}
}

func main() {
	stuffDoer := InitializeProcessor()
	stuffDoer.DoIO()
}
