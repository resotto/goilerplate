package testdata

import "github.com/resotto/goilerplate/cmd/app/domain/valueobject"

// MExchange is mock of service.IExchange
type MExchange struct{}

// Ticker is mock implementation of service.IExchange.Ticker()
func (e MExchange) Ticker(p valueobject.Pair) valueobject.Ticker {
	return valueobject.Ticker{
		Sell:      "1000",
		Buy:       "1000",
		High:      "2000",
		Low:       "500",
		Last:      "1200",
		Vol:       "20",
		Timestamp: "1600769562",
	}
}

// Ohlc is mock implementation of service.IExchange.Ohlc()
func (e MExchange) Ohlc(p valueobject.Pair, t valueobject.Timeunit) []valueobject.CandleStick {
	cs := make([]valueobject.CandleStick, 0)
	return append(cs, valueobject.CandleStick{
		Open:      "1000",
		High:      "2000",
		Low:       "500",
		Close:     "1500",
		Volume:    "30",
		Timestamp: "1600769562",
	})
}
