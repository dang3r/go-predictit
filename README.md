# go-predictit
A library exposing the API for the prediction market Predictit

## Installation
```sh
$ go get github.com/dang3r/go-predictit
```

## Example
--------------
```
package main

import (
  "fmt"
  "github.com/dang3r/gopredictit"
)

func main() {
  tickerSymbols := gopred.GetTickerSymbols()
  for i, v := range tickerSymbols {
    fmt.Println("%v", *gopred.GetMarketData(v))
  }
}
```
