package usecase

import (
	"time"

	"github.com/JayantiTA/backend-weights/internal/entity"
	"github.com/JayantiTA/backend-weights/internal/service"
)

type (
	Weight interface {
		GetAll() ([]entity.GetWeightsResponse, error)
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

func (u *weightUsecase) GetAll() ([]entity.GetWeightsResponse, error) {
	var response []entity.GetWeightsResponse
	weights, err := u.weightSvc.GetAll()
	if err != nil {
		return nil, err
	}
	for _, v := range weights {
		response = append(response, entity.GetWeightsResponse{
			ID:   v.ID,
			Date: v.Date,
			Max:  v.Max,
			Min:  v.Min,
			Diff: v.Max - v.Min,
		})
	}
	return response, nil
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
