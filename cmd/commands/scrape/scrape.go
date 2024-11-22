package scrape

import (
	"fmt"
	"log"
	"strings"

	"crypto-mine-cli/cmd/commands/saveResults"
	"crypto-mine-cli/config"

	"github.com/gocolly/colly"
	goPretty "github.com/jedib0t/go-pretty/v6/table"
)

func Scrape(saveResult bool, symbolFilter []string, fileType string) {
	if saveResult {
		if fileType != "json" && fileType != "csv" {
			log.Println("error: file type not supported, try json or csv")
		}

		if fileType == "json" {
			saveResults.SaveInJSON()
		}

		if fileType == "csv" {
			saveResults.SaveInCSV()
		}

		log.Printf("File persisted in the Downloads folder in %s format", fileType)
	}

	goPrettyTable := config.ConfigGoPretty()

	c := colly.NewCollector()

	c.OnHTML("tbody tr", func(h *colly.HTMLElement) {
		name := h.ChildText(fmt.Sprintf("%s--name", ClassColumn))
		symbol := h.ChildText(fmt.Sprintf("%s__symbol", ClassCell))
		marketCap := h.ChildText(fmt.Sprintf("%s__market-cap", ClassCell))
		price := h.ChildText(fmt.Sprintf("%s__price", ClassCell))
		volume := h.ChildText(fmt.Sprintf("%s__volume-24-h", ClassCell))
		change1h := h.ChildText(fmt.Sprintf("%s__percent-change-1-h", ClassCell))
		change24h := h.ChildText(fmt.Sprintf("%s__percent-change-24-h", ClassCell))
		change7d := h.ChildText(fmt.Sprintf("%s__percent-change-7-d", ClassCell))

		if len(symbolFilter) > 0 {
			for _, symbolValue := range symbolFilter {
				if symbolValue == strings.ToLower(symbol) {
					goPrettyTable.AppendRows([]goPretty.Row{
						{
							name,
							symbol,
							marketCap,
							price,
							volume,
							change1h,
							change24h,
							change7d,
						},
					})
				}
			}
			return
		}

		if name != "" {
			goPrettyTable.AppendRows([]goPretty.Row{
				{name, symbol, marketCap, price, volume, change1h, change24h, change7d},
			})
		}
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Fatalf("Something went wrong: %v", err)
	})

	c.Visit(CoinMarketURL)

	fmt.Println(goPrettyTable.Render())
}
