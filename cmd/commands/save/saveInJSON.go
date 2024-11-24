package save

import (
	"encoding/json"
	"fmt"
	"log"

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
		name := h.ChildText(fmt.Sprintf("%s--name", ClassColumn))
		symbol := h.ChildText(fmt.Sprintf("%s__symbol", ClassCell))
		marketCap := h.ChildText(fmt.Sprintf("%s__market-cap", ClassCell))
		price := h.ChildText(fmt.Sprintf("%s__price", ClassCell))
		volume := h.ChildText(fmt.Sprintf("%s__volume-24-h", ClassCell))
		change1h := h.ChildText(fmt.Sprintf("%s__percent-change-1-h", ClassCell))
		change24h := h.ChildText(fmt.Sprintf("%s__percent-change-24-h", ClassCell))
		change7d := h.ChildText(fmt.Sprintf("%s__percent-change-7-d", ClassCell))

		if name != "" {
			d := Data{
				Name:      name,
				Symbol:    symbol,
				MarketCap: marketCap,
				Price:     price,
				Volume:    volume,
				Change1h:  change1h,
				Change24h: change24h,
				Change7d:  change7d,
			}

			data = append(data, d)
		}
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Fatalf("Something went wrong: %v", err)
	})

	c.Visit(CoinMarketURL)

	fileCreated, _ := json.MarshalIndent(data, "", " ")

	if _, err = file.Write(fileCreated); err != nil {
		panic(err)
	}
}
