package usecase

import (
	"github.com/resotto/goilerplate/cmd/app/application/service"
	"github.com/resotto/goilerplate/cmd/app/domain/valueobject"
)

// Ticker is the usecase of getting ticker
func Ticker(e service.IExchange, p valueobject.Pair) valueobject.Ticker {
	return e.Ticker(p)
}
