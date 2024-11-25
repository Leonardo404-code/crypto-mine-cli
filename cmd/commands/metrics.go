package commands

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Metrics struct {
	Name      string
	Symbol    string
	MarketCap string
	Price     string
	Volume    string
	Change1h  string
	Change24h string
	Change7d  string
}

func GetMetrics(h *colly.HTMLElement) Metrics {
	name := h.ChildText(fmt.Sprintf("%s--name", ClassColumn))
	symbol := h.ChildText(fmt.Sprintf("%s__symbol", ClassCell))
	marketCap := h.ChildText(fmt.Sprintf("%s__market-cap", ClassCell))
	price := h.ChildText(fmt.Sprintf("%s__price", ClassCell))
	volume := h.ChildText(fmt.Sprintf("%s__volume-24-h", ClassCell))
	change1h := h.ChildText(fmt.Sprintf("%s__percent-change-1-h", ClassCell))
	change24h := h.ChildText(fmt.Sprintf("%s__percent-change-24-h", ClassCell))
	change7d := h.ChildText(fmt.Sprintf("%s__percent-change-7-d", ClassCell))

	metrics := Metrics{
		Name:      name,
		Symbol:    symbol,
		MarketCap: marketCap,
		Price:     price,
		Volume:    volume,
		Change1h:  change1h,
		Change24h: change24h,
		Change7d:  change7d,
	}

	return metrics
}
