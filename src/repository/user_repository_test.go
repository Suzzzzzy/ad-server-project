package repository

import (
	"ad-server-project/src/domain"
	"ad-server-project/src/domain/model"
	"context"
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type UserRepoTestSuite struct {
	suite.Suite
	userRepo model.UserRepository
	db       *sql.DB
	mock     sqlmock.Sqlmock
}

func (ts *UserRepoTestSuite) SetupTest() {
	var err error
	ts.db, ts.mock, err = sqlmock.New()
	if err != nil {
		fmt.Println("DB connection failed fot test:", err)
	}
	ts.userRepo = NewUserRepository(ts.db)
}

func TestUserRepoTestsuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}

var testUser = model.User{
	ID:        1,
	Gender:    "M",
	Country:   "KR",
	Reward:    100,
	CreatedAt: time.Now(),
}

func (ts *UserRepoTestSuite) Test_UpdateReward() {
	ts.Run("유저의 리워드 정보 업데이트", func() {
		// given
		ts.mock.ExpectPrepare("UPDATE users set reward = \\? WHERE id = \\?").
			ExpectExec().
			WithArgs(testUser.Reward, testUser.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		// when
		err := ts.userRepo.UpdateReward(context.Background(), testUser)

		if err != nil {
			ts.Fail("repository error", err)
		}
		// then
		if err := ts.mock.ExpectationsWereMet(); err != nil {
			ts.Fail("expectations were not met", err)
		}

	})
}

func (ts *UserRepoTestSuite) Test_GetById() {
	ts.Run("유저 고유값으로 유저정보 조회 - 성공", func() {
		// given
		rows := ts.mock.NewRows([]string{"id", "gender", "country", "reward", "created_at"}).AddRow(testUser.ID, testUser.Gender, testUser.Country, testUser.Reward, testUser.CreatedAt)
		ts.mock.ExpectQuery("SELECT (.+) FROM users WHERE id = \\?").
			WithArgs(testUser.ID).
			WillReturnRows(rows)

		// when
		result, err := ts.userRepo.GetById(context.Background(), testUser.ID)
		if err != nil {
			ts.Fail("repository error", err)
		}
		// then
		if err := ts.mock.ExpectationsWereMet(); err != nil {
			ts.Fail("expectations were not met", err)
		}
		ts.Equal(testUser, result)

	})
	ts.Run("유저 고유값으로 유저정보 조회 - 실패", func() {
		// given
		ts.mock.ExpectQuery("SELECT (.+) FROM users WHERE id = \\?").
			WithArgs(testUser.ID).
			WillReturnError(domain.ErrUserNotFound)
		// when
		_, err := ts.userRepo.GetById(context.Background(), testUser.ID)
		if err == nil {
			ts.Fail("repository error", err)
		}
		// then
		if err := ts.mock.ExpectationsWereMet(); err != nil {
			ts.Fail("expectations were not met", err)
		}
	})
}
