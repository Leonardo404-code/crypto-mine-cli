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

var SaveResults bool

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
		symbol := h.ChildText(fmt.Sprintf("%s__symbol", ClassText))
		marketCap := h.ChildText(fmt.Sprintf("%s__market-cap", ClassText))
		price := h.ChildText(fmt.Sprintf("%s__price", ClassText))
		volume := h.ChildText(fmt.Sprintf("%s__volume-24-h", ClassText))
		change1h := h.ChildText(fmt.Sprintf("%s__percent-change-1-h", ClassText))
		change24h := h.ChildText(fmt.Sprintf("%s__percent-change-24-h", ClassText))
		change7d := h.ChildText(fmt.Sprintf("%s__percent-change-7-d", ClassText))

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
