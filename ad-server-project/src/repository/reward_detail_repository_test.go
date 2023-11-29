package repository

import (
	"ad-server-project/src/domain/model"
	"context"
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/guregu/null"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type RewardDetailRepoTestSuite struct {
	suite.Suite
	rewardDetailRepo model.RewardDetailRepository
	db               *sql.DB
	mock             sqlmock.Sqlmock
}

func (ts *RewardDetailRepoTestSuite) SetupTest() {
	var err error
	ts.db, ts.mock, err = sqlmock.New()
	if err != nil {
		fmt.Println("DB connection failed fot test:", err)
	}
	ts.rewardDetailRepo = NewRewardDetailRepository(ts.db)
}

func TestRewardDetailRepoTestSuite(t *testing.T) {
	suite.Run(t, new(RewardDetailRepoTestSuite))
}

var testRewardDetail = []model.RewardDetail{
	{
		ID:        1,
		AdId:      null.IntFrom(1),
		UserId:    1,
		Reward:    10,
		CreatedAt: time.Now(),
	},
}

func (ts *RewardDetailRepoTestSuite) Test_EarnRewardDetail() {
	ts.Run("리워드 획득", func() {
		// given
		expectedResult := testRewardDetail[0]
		expectedResult.RewardType = model.Plus

		ts.mock.ExpectPrepare("INSERT reward_detail SET ad_id = \\?, reward = \\?, user_id = \\?, reward_type = \\?").
			ExpectExec().
			WithArgs(expectedResult.AdId, expectedResult.Reward, expectedResult.UserId, expectedResult.RewardType).
			WillReturnResult(sqlmock.NewResult(1, 1))

		// when
		err := ts.rewardDetailRepo.EarnRewardDetail(context.Background(), expectedResult.ID, expectedResult.Reward, expectedResult.UserId, expectedResult.RewardType)
		if err != nil {
			ts.Fail("repository error - %v", err)
		}
		// then
		if err := ts.mock.ExpectationsWereMet(); err != nil {
			ts.Fail("expectations were not met", err)
		}
	})
}

func (ts *RewardDetailRepoTestSuite) Test_DeductRewardDetail() {
	ts.Run("리워드 차감", func() {
		// given
		expectedResult := testRewardDetail[0]
		expectedResult.RewardType = model.Minus

		ts.mock.ExpectPrepare("INSERT reward_detail SET reward = \\?, user_id = \\?, reward_type = \\?").
			ExpectExec().
			WithArgs(expectedResult.Reward, expectedResult.UserId, expectedResult.RewardType).
			WillReturnResult(sqlmock.NewResult(1, 1))

		// when
		err := ts.rewardDetailRepo.DeductRewardDetail(context.Background(), expectedResult.Reward, expectedResult.UserId, expectedResult.RewardType)
		if err != nil {
			ts.Fail("repository error - %v", err)
		}
		// then
		if err := ts.mock.ExpectationsWereMet(); err != nil {
			ts.Fail("expectations were not met", err)
		}
	})
}

func (ts *RewardDetailRepoTestSuite) Test_GetRecent() {
	ts.Run("최근 리워드 내역 조회", func() {
		// given
		userId := 1
		expectedResult := testRewardDetail
		rows := ts.mock.NewRows([]string{"id", "ad_id", "user_id", "reward", "reward_type", "created_at"})
		for _, reward := range expectedResult {
			rows = rows.AddRow(reward.ID, reward.AdId, reward.UserId, reward.Reward, reward.RewardType, reward.CreatedAt)
		}

		ts.mock.ExpectQuery("SELECT (.+) FROM reward_detail WHERE user_id = \\? AND created_at >= NOW\\(\\) - INTERVAL 1 WEEK ORDER BY created_at DESC").
			WithArgs(userId).
			WillReturnRows(rows)
		// when
		result, err := ts.rewardDetailRepo.GetRecent(context.Background(), userId)
		if err != nil {
			ts.Fail("repository error - %v", err)
		}
		// then
		if err := ts.mock.ExpectationsWereMet(); err != nil {
			ts.Fail("expectations were not met", err)
		}
		ts.Equal(expectedResult, result)
	})
}

func (ts *RewardDetailRepoTestSuite) Test_GetAllRewardDetail() {
	ts.Run("최근 리워드 내역 조회", func() {
		// given
		userId := 1
		expectedResult := testRewardDetail
		rows := ts.mock.NewRows([]string{"id", "ad_id", "user_id", "reward", "reward_type", "created_at"})
		for _, reward := range expectedResult {
			rows = rows.AddRow(reward.ID, reward.AdId, reward.UserId, reward.Reward, reward.RewardType, reward.CreatedAt)
		}

		ts.mock.ExpectQuery("SELECT (.+) FROM reward_detail WHERE user_id = \\?").
			WithArgs(userId).
			WillReturnRows(rows)
		// when
		result, err := ts.rewardDetailRepo.GetAllRewardDetail(context.Background(), userId)
		if err != nil {
			ts.Fail("repository error - %v", err)
		}
		// then
		if err := ts.mock.ExpectationsWereMet(); err != nil {
			ts.Fail("expectations were not met", err)
		}
		ts.Equal(expectedResult, result)
	})
}
