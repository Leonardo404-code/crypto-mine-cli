package save

import (
	"log"

	"crypto-mine-cli/cmd/commands"

	"github.com/gocolly/colly"
)

func saveInCSV() {
	file, writer, err := createCSVFile()
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	defer writer.Flush()

	writer.Write(
		[]string{
			"Name",
			"Symbol",
			"Market capacity (USD)",
			"Price (USD)",
			"Volume (USD)",
			"Change (1h)",
			"Change (24h)",
			"Change (7d)",
		},
	)

	c := colly.NewCollector()

	c.OnHTML("tbody tr", func(h *colly.HTMLElement) {
		metrics := commands.GetMetrics(h)

		writer.Write([]string{
			metrics.Name,
			metrics.Symbol,
			metrics.MarketCap,
			metrics.Price,
			metrics.Volume,
			metrics.Change1h,
			metrics.Change24h,
			metrics.Change7d,
		})
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Fatalf("Something went wrong: %v", err)
	})

	c.Visit(commands.CoinMarketURL)
}
