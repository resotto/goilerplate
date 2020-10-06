package usecase

import (
	"github.com/resotto/goilerplate/internal/app/domain"
	"github.com/resotto/goilerplate/internal/app/domain/repository"
)

// Parameter is the usecase of getting parameter
func Parameter(r repository.IParameter) domain.Parameter {
	return r.Get()
}
