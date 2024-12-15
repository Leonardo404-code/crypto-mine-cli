package save

import (
	"encoding/json"
	"fmt"

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

func saveInJSON() error {
	file, err := createJSONFile()
	if err != nil {
		return err
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

	c.OnError(func(_ *colly.Response, requestErr error) {
		err = fmt.Errorf("Something went wrong: %v", requestErr)
	})

	if err != nil {
		return err
	}

	if err := c.Visit(commands.CoinMarketURL); err != nil {
		return fmt.Errorf("failed in visit %s", commands.CoinMarketURL)
	}

	fileCreated, _ := json.MarshalIndent(data, "", " ")

	if _, err = file.Write(fileCreated); err != nil {
		return fmt.Errorf("failed in write file data: %v", err)
	}

	return nil
}
