package scrape

import (
	"fmt"
	"strings"

	"crypto-mine-cli/cmd/commands"
	"crypto-mine-cli/config"

	"github.com/gocolly/colly"
	goPretty "github.com/jedib0t/go-pretty/v6/table"
)

func Scrape(symbolFilter []string) error {
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

	var err error

	c.OnError(func(_ *colly.Response, requestErr error) {
		err = fmt.Errorf("Something went wrong: %v", requestErr)
	})

	if err != nil {
		return err
	}

	if err = c.Visit(commands.CoinMarketURL); err != nil {
		return err
	}

	fmt.Println(goPrettyTable.Render())

	return nil
}
