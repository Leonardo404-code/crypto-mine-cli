package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "crypto-mine",
	Short: "crypto-mine-cli application on CoinMarketCap website to get information of all cryptocurrencies",
	Long: `Crypto crypto-mine-cli is an application that extracts the
data of various cryptocurrencies from the Coin Market Cap
	`,
}

func init() {
	scrapeCmd.Flags().BoolVarP(
		&SaveResults,
		"save-result",
		"s",
		false,
		"Save result stores the results in a CSV file in the Downloads folder",
	)
	rootCmd.AddCommand(scrapeCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
