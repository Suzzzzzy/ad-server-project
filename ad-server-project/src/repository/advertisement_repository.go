package repository

import (
	"ad-server-project/src/domain"
	"ad-server-project/src/domain/model"
	"context"
	"database/sql"
	"fmt"
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
	query := `SELECT * FROM advertisement WHERE target_gender = ? AND target_country = ?`
	list, err := a.fetch(c, query, user.Gender, user.Country)
	if err != nil {
		log.Printf("advertisementRepository query Error: %v \n", err)
		return nil, err
	}

	return list, nil
}

func (a *advertisementRepository) UpdateReward(c context.Context, id int, reward int) error {
	query := `UPDATE advertisement set reward = ? WHERE id = ?`

	stmt, err := a.Conn.PrepareContext(c, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(c, reward, id)
	if err != nil {
		return err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect != 1 {
		return fmt.Errorf("No change in reward value")
	}
	return nil
}

func (a *advertisementRepository) GetById(c context.Context, id int) (result model.Advertisement, err error) {
	query := `SELECT * FROM advertisement WHERE id = ?`
	res, err := a.fetch(c, query, id)
	if err != nil {
		return model.Advertisement{}, err
	}
	if len(res) > 0 {
		result = res[0]
	} else {
		return result, domain.ErrNotFound
	}
	return result, err
}
