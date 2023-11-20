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

func (a *advertisementRepository) GetByCountryAndGender(c context.Context, user *model.User) ([]model.Advertisement, error) {
	//TODO implement me
	panic("implement me")
}
