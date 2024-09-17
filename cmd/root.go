package cmd

import (
	"os"

	"crypto-mine-cli/pkg/scrape"

	"github.com/spf13/cobra"
)

var SaveResults bool

var rootCmd = &cobra.Command{
	Use:   "crypto-mine",
	Short: "crypto-mine-cli collects cryptocurrency data on the coin market cape website",
	Long: `Crypto crypto-mine-cli is an application that extracts the
data of various cryptocurrencies from the Coin Market Cap
	`,
	Run: func(cmd *cobra.Command, args []string) {
		scrape.Scrape()
	},
}

func init() {
	rootCmd.Flags().BoolVarP(
		&SaveResults,
		"save-result",
		"s",
		false,
		"Save result stores the results in a CSV file in the Downloads folder",
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
