package cmd

import (
	"crypto-mine-cli/cmd/commands/compare"
	"crypto-mine-cli/cmd/commands/save"
	"crypto-mine-cli/cmd/commands/scrape"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "crypto-mine",
		Short: "crypto-mine-cli collects cryptocurrency data on the coin market cape website",
		Long: `Crypto crypto-mine-cli is an application that extracts the
data of various cryptocurrencies from the Coin Market Cap
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := scrape.Scrape(filter); err != nil {
				return err
			}

			return nil
		},
	}

	saveCmd = &cobra.Command{
		Use:   "save",
		Short: "Save stores the results in Donwloads folder",
		Long:  "Save stores the results in Downloads folder, The information from the table is applied to the file of your choice, choose between JSON or CSV with the --type flag (or -t for short)",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := save.Save(fileType); err != nil {
				return err
			}

			return nil
		},
	}

	compareCmd = &cobra.Command{
		Use:   "compare",
		Short: "Short description",
		Long:  "Long description",
		Run: func(cmd *cobra.Command, args []string) {
			compare.Compare(cryptos, metrics)
		},
	}
)
