package gopred

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// The routes for all markets
var routeSet = make(map[string]struct{})

// The routes for all individual contracts
var contractSet = make(map[string]struct{})

// An entry representing an event
type MarketEntry struct {
	Route        string
	TickerSymbol string
	Data         Result
}

// Scrape the predictit website to extract all of the present ticker symbols
func GetTickerSymbols() []string {
	var tickerSymbols []string
	purl := url.URL{
		Scheme: "https",
		Host:   "www.predictit.org",
		Path:   "Browse/Featured",
	}

	// Retrieve homepage
	doc, err := goquery.NewDocument(purl.String())
	if err != nil {
		log.Fatalf("Error parsing predictit homepage : %s", err)
	}

	// Retrieve all market subpages
	// The goqueries retrieve duplicates so use a set
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		if s.HasClass("dropdown-toggle") {
			if val, exists := s.Attr("href"); exists && strings.Index(val, "Group") != -1 {
				routeSet[val] = struct{}{}
			}
		}
	})

	// For each market subpage, determine all events on each
	for subpage, _ := range routeSet {
		purl.Path = subpage
		doc, err := goquery.NewDocument(purl.String())
		if err != nil {
			log.Fatalf("Unable to parse market subpage : %s", err)
		}

		doc.Find("a").Each(func(i int, s *goquery.Selection) {
			if val, exists := s.Attr("href"); exists && strings.Index(val, "Contract") != -1 {
				contractSet[val] = struct{}{}
			}
		})
	}

	// For each event, extract the ticker-symbol
	for path, _ := range contractSet {
		purl.Path = path
		doc, _ := goquery.NewDocument(purl.String())
		if err != nil {
			log.Fatalf("Unable to parse contract document for %s : %s", path, err)
		}

		doc.Find("table").Each(func(i int, s *goquery.Selection) {
			if s.HasClass("table table-condensed table-striped table-info") {
				s.ChildrenFiltered("tbody").ChildrenFiltered("tr").First().Children().Next().Each(func(i int, s *goquery.Selection) {
					tickerSymbols = append(tickerSymbols, s.Text())
				})
			}
		})
	}
	return tickerSymbols
}

// Retrieve all exposed marketdata for a tickerSymbol
// https://predictit.freshdesk.com/support/solutions/articles/
// 	12000001878-does-predictit-make-market-data-available-via-an-api-
func GetMarketData(tickerSymbol string) *Result {
	marketData := new(Result)
	purl := url.URL{
		Scheme: "https",
		Host:   "www.predictit.org",
		Path:   "/api/marketdata/ticker/" + tickerSymbol,
	}
	res, err := http.Get(purl.String())
	defer res.Body.Close()
	if err != nil {
		return nil
	}
	bytes, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(bytes, marketData)
	if err != nil {
		log.Fatalf("Unable to unmarshal : %s", err)
	}

	return marketData
}
