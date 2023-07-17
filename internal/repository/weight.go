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

func (r *weightRepo) CreateOrUpdate(weight entity.WeightDto) error {
	weightDao := entity.Weight{
		Date: weight.Date,
		Max:  weight.Max,
		Min:  weight.Min,
	}
	weightResult := r.db.Where(entity.Weight{Date: weight.Date}).FirstOrCreate(&weightDao)
	if weightResult.Error != nil {
		return weightResult.Error
	}

	if weightResult.RowsAffected == 0 {
		return r.db.Save(&weight).Error
	}

	return nil
}

func (r *weightRepo) Delete(date *time.Time) error {
	if err := r.db.Delete(&entity.Weight{Date: date}).Error; err != nil {
		return err
	}
	return nil
}
