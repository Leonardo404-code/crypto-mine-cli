package save

import (
	"encoding/json"
	"log"

	"crypto-mine-cli/cmd/commands"

	"github.com/gocolly/colly"
)

type Data struct {
	Name      string `json:"name"`
	Symbol    string `json:"symbol"`
	MarketCap string `json:"market_cap"`
	Price     string `json:"price"`
	Volume    string `json:"volume"`
	Change1h  string `json:"change1h"`
	Change24h string `json:"change24h"`
	Change7d  string `json:"change7d"`
}

func saveInJSON() {
	file, err := createJSONFile()
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	c := colly.NewCollector()

	data := []Data{}

	c.OnHTML("tbody tr", func(h *colly.HTMLElement) {
		metrics := commands.GetMetrics(h)

		if metrics.Name != "" {
			d := Data{
				Name:      metrics.Name,
				Symbol:    metrics.Symbol,
				MarketCap: metrics.MarketCap,
				Price:     metrics.Price,
				Volume:    metrics.Volume,
				Change1h:  metrics.Change1h,
				Change24h: metrics.Change24h,
				Change7d:  metrics.Change7d,
			}

			data = append(data, d)
		}
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Fatalf("Something went wrong: %v", err)
	})

	c.Visit(commands.CoinMarketURL)

	fileCreated, _ := json.MarshalIndent(data, "", " ")

	if _, err = file.Write(fileCreated); err != nil {
		panic(err)
	}
}
