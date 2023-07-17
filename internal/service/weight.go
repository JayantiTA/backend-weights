package service

import (
	"time"

	"github.com/JayantiTA/backend-weights/internal/entity"
	"github.com/JayantiTA/backend-weights/internal/repository"
)

type (
	Weight interface {
		GetAll() ([]entity.Weight, error)
		Get(date *time.Time) (entity.Weight, error)
		CreateOrUpdate(weight entity.WeightDto) error
		Delete(date *time.Time) error
	}

	weightSvc struct {
		weightRepo repository.Weight
	}
)

func NewWeight(weightRepo repository.Weight) Weight {
	return &weightSvc{
		weightRepo: weightRepo,
	}
}

func (s *weightSvc) GetAll() ([]entity.Weight, error) {
	return s.weightRepo.GetAll()
}

func (s *weightSvc) Get(date *time.Time) (entity.Weight, error) {
	return s.weightRepo.Get(date)
}

func (s *weightSvc) CreateOrUpdate(weight entity.WeightDto) error {
	return s.weightRepo.CreateOrUpdate(weight)
}

func (s *weightSvc) Delete(date *time.Time) error {
	return s.weightRepo.Delete(date)
}
