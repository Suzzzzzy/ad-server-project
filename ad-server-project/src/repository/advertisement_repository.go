package repository

import (
	"ad-server-project/src/domain/model"
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
	"log"
)

type advertisementRepository struct {
	Conn *sql.DB
}

func NewAdvertisementRepository(conn *sql.DB) model.AdvertisementRepository {
	return &advertisementRepository{
		Conn: conn,
	}
}

func (a *advertisementRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []model.Advertisement, err error) {
	rows, err := a.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]model.Advertisement, 0)
	for rows.Next() {
		t := model.Advertisement{}
		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.ImageUrl,
			&t.LandingUrl,
			&t.Weight,
			&t.TargetCountry,
			&t.TargetGender,
			&t.Reward,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}
	return result, nil
}

func (a *advertisementRepository) GetByCountryAndGender(c context.Context, user *model.User) ([]model.Advertisement, error) {
	query := `SELECT * FROM advertisement WHERE target_gender = ? and target_country = ?`
	list, err := a.fetch(c, query, user.Gender, user.Country)
	if err != nil {
		log.Printf("advertisementRepository query Error: %v \n", err)
		return nil, err
	}

	return list, nil
}
