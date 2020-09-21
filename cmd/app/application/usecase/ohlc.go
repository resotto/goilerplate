package usecase

import (
	"github.com/resotto/goilerplate/cmd/app/application/service"
	"github.com/resotto/goilerplate/cmd/app/domain/valueobject"
)

// Ohlc is the usecase of getting open, high, low, and close
func Ohlc(e service.IExchange, p valueobject.Pair, t valueobject.Timeunit) []valueobject.CandleStick {
	return e.Ohlc(p, t)
}
