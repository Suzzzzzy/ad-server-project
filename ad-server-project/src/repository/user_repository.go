package repository

import (
	"ad-server-project/src/domain"
	"ad-server-project/src/domain/model"
	"context"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
)

type userRepository struct {
	Conn *sql.DB
}

func NewUserRepository(conn *sql.DB) model.UserRepository {
	return &userRepository{
		Conn: conn,
	}
}

func (u *userRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []model.User, err error) {
	rows, err := u.Conn.QueryContext(ctx, query, args...)
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

	result = make([]model.User, 0)
	for rows.Next() {
		t := model.User{}
		err = rows.Scan(
			&t.ID,
			&t.Gender,
			&t.Country,
			&t.Reward,
			&t.CreatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}
	return result, nil
}

func (u *userRepository) UpdateReward(c context.Context, user model.User) error {
	query := `UPDATE users set reward = ? WHERE id = ?`

	stmt, err := u.Conn.PrepareContext(c, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(c, user.Reward, user.ID)
	if err != nil {
		return err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect != 1 {
		return fmt.Errorf("[Error] No row was affected for reward with user_id: %v", user.ID)
	}
	return nil
}

func (u *userRepository) GetById(c context.Context, userId int) (result model.User, err error) {
	query := `SELECT * FROM users WHERE id = ?`

	res, err := u.fetch(c, query, userId)
	if err != nil {
		return model.User{}, err
	}
	if len(res) > 0 {
		result = res[0]
	} else {
		return result, domain.ErrUserNotFound
	}
	return result, err
}