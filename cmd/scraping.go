/*
Copyright © 2024 NAME HERE <leonardobispo.dev@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
	goPretty "github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Scrape collects cryptocurrency data on the coin market cape website",
	Long: `Scrape Collects cryptocurrency data and stores it directly in the downloads folder,
being able to specify a specific currency`,
	Run: func(cmd *cobra.Command, args []string) {
		Scrape()
	},
}

func configGoPretty() goPretty.Writer {
	tw := goPretty.NewWriter()
	tw.AppendHeader(goPretty.Row{
		"Name",
		"Symbol",
		"Market capacity (USD)",
		"Price (USD)",
		"Volume (USD)",
		"Change (1h)",
		"Change (24h)",
		"Change (7d)",
	})

	return tw
}

func Scrape() {
	goPrettyTable := configGoPretty()

	c := colly.NewCollector()

	c.OnHTML("tbody tr", func(h *colly.HTMLElement) {
		name := h.ChildText(".cmc-table__column-name--name")
		symbol := h.ChildText(".cmc-table__cell--sort-by__symbol")
		marketCap := h.ChildText(".cmc-table__cell--sort-by__market-cap")
		price := h.ChildText(".cmc-table__cell--sort-by__price")
		volume := h.ChildText(".cmc-table__cell--sort-by__volume-24-h")
		change1h := h.ChildText(".cmc-table__cell--sort-by__percent-change-1-h")
		change24h := h.ChildText(".cmc-table__cell--sort-by__percent-change-24-h")
		change7d := h.ChildText(".cmc-table__cell--sort-by__percent-change-7-d")

		if name != "" {
			goPrettyTable.AppendRows([]goPretty.Row{
				{name, symbol, marketCap, price, volume, change1h, change24h, change7d},
			})
		}
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Fatalf("Something went wrong: %v", err)
	})

	c.Visit("https://coinmarketcap.com/all/views/all/")

	fmt.Println(goPrettyTable.Render())
}
