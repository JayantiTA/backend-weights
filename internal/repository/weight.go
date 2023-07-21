package repository

import (
	"time"

	"github.com/JayantiTA/backend-weights/internal/entity"
	"gorm.io/gorm"
)

type (
	Weight interface {
		GetAll() ([]entity.Weight, error)
		Get(date *time.Time) (entity.Weight, error)
		CreateOrUpdate(weight entity.WeightDto) error
		Delete(date *time.Time) error
	}

	weightRepo struct {
		db *gorm.DB
	}
)

func NewWeight(db *gorm.DB) Weight {
	return &weightRepo{
		db: db,
	}
}

func (r *weightRepo) GetAll() ([]entity.Weight, error) {
	var weights []entity.Weight
	if err := r.db.Find(&weights).Error; err != nil {
		return nil, err
	}
	return weights, nil
}

func (r *weightRepo) Get(date *time.Time) (entity.Weight, error) {
	var weight entity.Weight
	if err := r.db.Model(entity.Weight{Date: date}).First(&weight).Error; err != nil {
		return entity.Weight{}, err
	}
	return weight, nil
}

func (r *weightRepo) CreateOrUpdate(weightRequest entity.WeightDto) error {
	weightDao := entity.Weight{
		Date: weightRequest.Date,
		Max:  weightRequest.Max,
		Min:  weightRequest.Min,
	}
	weight := r.db.Where(entity.Weight{Date: weightRequest.Date}).FirstOrCreate(&weightDao)
	if weight.Error != nil {
		return weight.Error
	}

	if weight.RowsAffected == 0 {
		return r.db.Model(entity.Weight{}).Where("date = ?", weightRequest.Date).Updates(entity.Weight{Max: weightRequest.Max, Min: weightRequest.Min}).Error
	}

	return nil
}

func (r *weightRepo) Delete(date *time.Time) error {
	return r.db.Where("date = ?", date).Delete(&entity.Weight{}).Error
}
