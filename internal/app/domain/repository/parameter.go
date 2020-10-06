package repository

import "github.com/resotto/goilerplate/internal/app/domain"

// IParameter is interface of parameter repository
type IParameter interface {
	Get() domain.Parameter
}
