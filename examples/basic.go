// A simple program to demonstate the two main function calls of the exposed API
package main

import (
	gopred "github.com/dang3r/go-predictit"
	"log"
)

func main() {
	symbols := gopred.GetTickerSymbols()
	log.Println(len(symbols), "retrieved")
	if len(symbols) > 0 {
		result := gopred.GetMarketData(symbols[0])
		log.Println(result)
	} else {
		log.Println("No ticker symbols retrieved :(")
	}
}
