package usecase

import (
	"github.com/resotto/goilerplate/internal/app/application/service"
	"github.com/resotto/goilerplate/internal/app/domain/valueobject"
)

// Ticker is the usecase of getting ticker
func Ticker(e service.IExchange, p valueobject.Pair) valueobject.Ticker {
	return e.Ticker(p)
}
