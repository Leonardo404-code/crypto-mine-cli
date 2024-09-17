package saveResults

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func SaveInFile() {
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
		name := h.ChildText(fmt.Sprintf("%s--name", ClassColumn))
		symbol := h.ChildText(fmt.Sprintf("%s__symbol", ClassCell))
		marketCap := h.ChildText(fmt.Sprintf("%s__market-cap", ClassCell))
		price := h.ChildText(fmt.Sprintf("%s__price", ClassCell))
		volume := h.ChildText(fmt.Sprintf("%s__volume-24-h", ClassCell))
		change1h := h.ChildText(fmt.Sprintf("%s__percent-change-1-h", ClassCell))
		change24h := h.ChildText(fmt.Sprintf("%s__percent-change-24-h", ClassCell))
		change7d := h.ChildText(fmt.Sprintf("%s__percent-change-7-d", ClassCell))

		writer.Write([]string{
			name,
			symbol,
			marketCap,
			price,
			volume,
			change1h,
			change24h,
			change7d,
		})
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Fatalf("Something went wrong: %v", err)
	})

	c.Visit(CoinMarketURL)
}
