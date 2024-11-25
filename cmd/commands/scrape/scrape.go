package scrape

import (
	"fmt"
	"log"
	"strings"

	"crypto-mine-cli/cmd/commands"
	"crypto-mine-cli/config"

	"github.com/gocolly/colly"
	goPretty "github.com/jedib0t/go-pretty/v6/table"
)

func Scrape(symbolFilter []string) {
	goPrettyTable := config.ConfigGoPretty()

	c := colly.NewCollector()

	c.OnHTML("tbody tr", func(h *colly.HTMLElement) {
		metrics := commands.GetMetrics(h)

		if len(symbolFilter) > 0 {
			for _, symbolValue := range symbolFilter {
				if symbolValue == strings.ToLower(metrics.Symbol) {
					goPrettyTable.AppendRows([]goPretty.Row{
						{
							metrics.Name,
							metrics.Symbol,
							metrics.MarketCap,
							metrics.Price,
							metrics.Volume,
							metrics.Change1h,
							metrics.Change24h,
							metrics.Change7d,
						},
					})
				}
			}
			return
		}

		if metrics.Symbol != "" {
			goPrettyTable.AppendRows([]goPretty.Row{
				{
					metrics.Name,
					metrics.Symbol,
					metrics.MarketCap,
					metrics.Price,
					metrics.Volume,
					metrics.Change1h,
					metrics.Change24h,
					metrics.Change7d,
				},
			})
		}
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Fatalf("Something went wrong: %v", err)
	})

	c.Visit(commands.CoinMarketURL)

	fmt.Println(goPrettyTable.Render())
}
