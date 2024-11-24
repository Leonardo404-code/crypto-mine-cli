package cmd

import (
	"os"
)

func init() {
	rootCmd.Flags().StringArrayVarP(
		&symbolFilter,
		"filter",
		"f",
		[]string{},
		"Filter Crypto by symbol",
	)

	saveCmd.Flags().StringVarP(
		&fileType,
		"type",
		"t",
		"csv",
		"Determine the type of file that will be created to persist the cryptocurrency data, a CSV or JSON file.",
	)

	rootCmd.AddCommand(saveCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
