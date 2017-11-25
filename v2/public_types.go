package bitfinex

import (
	"fmt"
)

type Ticker struct {
	Symbol          string
	Bid             float64
	BidPeriod       int64
	BidSize         float64
	Ask             float64
	AskPeriod       int64
	AskSize         float64
	DailyChange     float64
	DailyChangePerc float64
	LastPrice       float64
	Volume          float64
	High            float64
	Low             float64
}

func (t *Ticker) String() string{
	return fmt.Sprintf("Symbol:%v, LastPrice:%v", t.Symbol, t.LastPrice)

}

type TickerUpdate Ticker
type TickerSnapshot []Ticker
