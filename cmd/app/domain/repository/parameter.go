package repository

import "github.com/resotto/goilerplate/cmd/app/domain"

// IParameter is interface of parameter repository
type IParameter interface {
	Get() domain.Parameter
	Save(domain.Parameter)
}
