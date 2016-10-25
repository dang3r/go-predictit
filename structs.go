package gopred

// The set of all

type Contract struct {
	Id              int     `json:"ID"`
	DateEnd         string  `json:"DateEnd"`
	Image           string  `json:"Image"`
	URL             string  `json:"URL"`
	Name            string  `json:"Name"`
	ShortName       string  `json:"ShortName"`
	LongName        string  `json:"LongName"`
	TickerSymbol    string  `json:"TickerSymbol"`
	Status          string  `json:"Status"`
	LastTradePrice  float64 `json:"LastTradePrice"`
	BestBuyYesCost  float64 `json:"BestBuyYesCost"`
	BestBuyNoCost   float64 `json:"BestBuyNoCost"`
	BestSellYesCost float64 `json:"BestSellYesCost"`
	BestSellNoCost  float64 `json:"BestSellNoCost"`
	LastClosePrice  float64 `json:"LastClosePrice"`
}

type Result struct {
	Id           int        `json:"ID"`
	Name         string     `json:"Name"`
	Shortname    string     `json:"ShortName"`
	TickerSymbol string     `json:"TickerSymbol"`
	Image        string     `json:"Image"`
	URL          string     `json:"URL"`
	Contracts    []Contract `json:"Contracts"`
	TimeStamp    string     `json:"TimeStamp"`
	Status       string     `json:"Open"`
}
