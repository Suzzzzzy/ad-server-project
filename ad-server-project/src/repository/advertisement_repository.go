package repository

import (
	"ad-server-project/src/domain/model"
	"context"
	"gorm.io/gorm"
)

type advertisementRepository struct {
	db *gorm.DB
}

func NewAdvertisementRepository(db *gorm.DB) model.AdvertisementRepository {
	return &advertisementRepository{
		db: db,
	}
}

func (a *advertisementRepository) GetByCountryAndGender(c context.Context, user *model.User) ([]map[string]interface{}, error) {
	var result []map[string]interface{}

	err := a.db.WithContext(c).Model(&model.Advertisement{}).
		Where("target_gender = ? and target_country = ?", user.Gender, user.Country).Find(result).Error

	return result, err
}
