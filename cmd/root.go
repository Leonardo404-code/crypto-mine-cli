package cmd

import (
	"os"

	"crypto-mine-cli/cmd/commands/scrape"

	"github.com/spf13/cobra"
)

var (
	saveResults  bool
	symbolFilter []string
	fileType     string
)

var rootCmd = &cobra.Command{
	Use:   "crypto-mine",
	Short: "crypto-mine-cli collects cryptocurrency data on the coin market cape website",
	Long: `Crypto crypto-mine-cli is an application that extracts the
data of various cryptocurrencies from the Coin Market Cap
	`,
	Run: func(cmd *cobra.Command, args []string) {
		scrape.Scrape(saveResults, symbolFilter, fileType)
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

	rootCmd.Flags().StringVarP(
		&fileType,
		"type",
		"t",
		"csv",
		"Determine the type of file that will be created to persist the cryptocurrency data, a CSV or JSON file.")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
