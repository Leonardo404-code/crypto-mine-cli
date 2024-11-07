package config

import (
	goPretty "github.com/jedib0t/go-pretty/v6/table"
)

func ConfigGoPretty() goPretty.Writer {
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
