package cmd

import (
	"os"

	"crypto-mine-cli/cmd/commands/scrape"

	"github.com/spf13/cobra"
)

var (
	saveResults  bool
	symbolFilter []string
)

var rootCmd = &cobra.Command{
	Use:   "crypto-mine",
	Short: "crypto-mine-cli collects cryptocurrency data on the coin market cape website",
	Long: `Crypto crypto-mine-cli is an application that extracts the
data of various cryptocurrencies from the Coin Market Cap
	`,
	Run: func(cmd *cobra.Command, args []string) {
		scrape.Scrape(saveResults, symbolFilter)
	},
}

func init() {
	rootCmd.Flags().BoolVarP(
		&saveResults,
		"save",
		"s",
		false,
		"Save result stores the results in a CSV file in the Downloads folder",
	)

	rootCmd.Flags().StringArrayVarP(
		&symbolFilter,
		"filter",
		"f",
		[]string{},
		"Filter Crypto by symbol",
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
