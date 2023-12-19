package scraping

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func InitScraping() {
	file, err := os.Create(FileName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", FileName, err)
		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	scraping(writer)
}

func scraping(writer *csv.Writer) {
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
		writer.Write([]string{
			h.ChildText(".cmc-table__column-name--name"),
			h.ChildText(".cmc-table__cell--sort-by__symbol"),
			h.ChildText(".cmc-table__cell--sort-by__market-cap"),
			h.ChildText(".cmc-table__cell--sort-by__price"),
			h.ChildText(".cmc-table__cell--sort-by__volume-24-h"),
			h.ChildText(".cmc-table__cell--sort-by__percent-change-1-h"),
			h.ChildText(".cmc-table__cell--sort-by__percent-change-24-h"),
			h.ChildText(".cmc-table__cell--sort-by__percent-change-7-d"),
		})
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Fatalf("Something went wrong: %v", err)
	})

	c.Visit("https://coinmarketcap.com/all/views/all/")
}
