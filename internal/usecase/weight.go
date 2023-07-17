package usecase

import (
	"time"

	"github.com/JayantiTA/backend-weights/internal/entity"
	"github.com/JayantiTA/backend-weights/internal/service"
)

type (
	Weight interface {
		GetAll() ([]entity.Weight, error)
		Get(date *time.Time) (entity.Weight, error)
		CreateOrUpdate(date *time.Time, max, min int) error
		Delete(date *time.Time) error
	}

	weightUsecase struct {
		weightSvc service.Weight
	}
)

func NewWeight(weightSvc service.Weight) Weight {
	return &weightUsecase{
		weightSvc: weightSvc,
	}
}

func (u *weightUsecase) GetAll() ([]entity.Weight, error) {
	return u.weightSvc.GetAll()
}

func (u *weightUsecase) Get(date *time.Time) (entity.Weight, error) {
	return u.weightSvc.Get(date)
}

func (u *weightUsecase) CreateOrUpdate(date *time.Time, max, min int) error {
	weight := entity.WeightDto{
		Date: date,
		Max:  max,
		Min:  min,
	}
	return u.weightSvc.CreateOrUpdate(weight)
}

func (u *weightUsecase) Delete(date *time.Time) error {
	return u.weightSvc.Delete(date)
}
