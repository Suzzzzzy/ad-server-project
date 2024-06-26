package repository

import (
	"ad-server-project/src/domain/model"
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
	"time"
)

type rewardDetailRepository struct {
	Conn *sql.DB
}

func NewRewardDetailRepository(conn *sql.DB) model.RewardDetailRepository {
	return &rewardDetailRepository{
		Conn: conn,
	}
}

func (r *rewardDetailRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []model.RewardDetail, err error) {
	rows, err := r.Conn.QueryContext(ctx, query, args...)
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

	var createdAt time.Time
	result = make([]model.RewardDetail, 0)
	for rows.Next() {
		t := model.RewardDetail{}
		if err = rows.Scan(
			&t.ID,
			&t.AdId,
			&t.UserId,
			&t.Reward,
			&t.RewardType,
			&createdAt,
		); err != nil {
			logrus.Error(err)
			return nil, err
		}
		t.CreatedAt = createdAt
		result = append(result, t)
	}

	if err = rows.Err(); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return result, nil
}

func (r *rewardDetailRepository) EarnRewardDetail(c context.Context, adId int, reward int, userId int, rewardType string) error {
	query := `INSERT reward_detail SET ad_id = ?, reward = ?, user_id = ?, reward_type = ? `
	stmt, err := r.Conn.PrepareContext(c, query)
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(c, adId, reward, userId, rewardType)
	if err != nil {
		return err
	}

	return nil
}

func (r *rewardDetailRepository) DeductRewardDetail(c context.Context, reward int, userId int, rewardType string) error {
	query := `INSERT reward_detail SET reward = ?, user_id = ?, reward_type = ? `
	stmt, err := r.Conn.PrepareContext(c, query)
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(c, reward, userId, rewardType)
	if err != nil {
		return err
	}

	return nil
}

func (r *rewardDetailRepository) GetRecent(c context.Context, userId int) (res []model.RewardDetail, err error) {
	query := `SELECT * FROM reward_detail WHERE user_id = ? AND created_at >= NOW() - INTERVAL 1 WEEK ORDER BY created_at DESC`
	res, err = r.fetch(c, query, userId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *rewardDetailRepository) GetAllRewardDetail(c context.Context, userId int) (res []model.RewardDetail, err error) {
	query := `SELECT * FROM reward_detail WHERE user_id = ?`
	res, err = r.fetch(c, query, userId)
	if err != nil {
		return nil, err
	}

	return res, nil
}
