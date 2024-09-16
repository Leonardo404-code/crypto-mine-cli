package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "web-scraping-colly",
	Short: "Web scraping application on CoinMarketCap website to get information of all cryptocurrencies",
	Long: `Crypto Web Scraping is an application that extracts the
data of various cryptocurrencies from the Coin Market Cap
	`,
}

func init() {
	rootCmd.AddCommand(scrapeCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
