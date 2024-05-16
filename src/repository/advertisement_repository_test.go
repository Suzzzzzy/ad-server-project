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
)

type AdvertisementRepoTestSuite struct {
	suite.Suite
	advertisementRepo model.AdvertisementRepository
	db                *sql.DB
	mock              sqlmock.Sqlmock
}

func (ts *AdvertisementRepoTestSuite) SetupTest() {
	var err error
	ts.db, ts.mock, err = sqlmock.New()
	if err != nil {
		fmt.Println("DB connection failed fot test:", err)
	}
	ts.advertisementRepo = NewAdvertisementRepository(ts.db)
}

func TestAdvertisementRepoTestsuite(t *testing.T) {
	suite.Run(t, new(AdvertisementRepoTestSuite))
}

var testAdvertisement = []model.Advertisement{
	{
		ID:            1,
		Name:          "Ad1",
		ImageUrl:      "http://example.com/ad1.jpg",
		LandingUrl:    "http://example.com/landing1",
		Weight:        1,
		TargetCountry: "USA",
		TargetGender:  "M",
		Reward:        100,
	},
}

func (ts *AdvertisementRepoTestSuite) Test_GetByCountryAndGender() {
	ts.Run("국가, 성별로 광고 조회", func() {
		// given
		testUser := &model.User{Country: "USA", Gender: "M"}
		expectedResult := testAdvertisement

		rows := ts.mock.NewRows([]string{"id", "name", "image_url", "landing_url", "weight", "target_country", "target_gender", "reward"})

		for _, ad := range expectedResult {
			rows = rows.AddRow(ad.ID, ad.Name, ad.ImageUrl, ad.LandingUrl, ad.Weight, ad.TargetCountry, ad.TargetGender, ad.Reward)
		}
		ts.mock.ExpectQuery("SELECT (.+) FROM advertisement WHERE target_gender = \\? AND target_country = \\?").
			WithArgs(testUser.Gender, testUser.Country).
			WillReturnRows(rows)

		// when
		result, err := ts.advertisementRepo.GetByCountryAndGender(context.Background(), testUser)
		if err != nil {
			ts.Fail("repository error", err)
		}
		// then
		ts.Equal(1, len(result))
		ts.Equal(expectedResult[0], result[0])

	})
}

func (ts *AdvertisementRepoTestSuite) Test_UpdateReward() {
	ts.Run("리워드 업데이트 - 성공", func() {
		// given
		id := 1
		reward := 100
		ts.mock.ExpectPrepare("UPDATE advertisement set reward = \\? WHERE id = \\?").
			ExpectExec().
			WithArgs(reward, id).
			WillReturnResult(sqlmock.NewResult(1, 1))
		// when
		err := ts.advertisementRepo.UpdateReward(context.Background(), id, reward)
		// then
		if err != nil {
			ts.Fail("repository error", err)
		}
		if err := ts.mock.ExpectationsWereMet(); err != nil {
			ts.Fail("expectations were not met", err)
		}

	})
	ts.Run("리워드 업데이트 - 실패", func() {
		// given
		id := 1
		reward := 100
		ts.mock.ExpectPrepare("UPDATE advertisement set reward = \\? WHERE id = \\?").
			ExpectExec().
			WithArgs(reward, id).
			WillReturnError(domain.ErrNotFound)
		// when
		err := ts.advertisementRepo.UpdateReward(context.Background(), id, reward)
		// then
		if err == nil {
			ts.Fail("expected prepare error, got", err)
		}
		if err := ts.mock.ExpectationsWereMet(); err != nil {
			ts.Fail("expectations were not met", err)
		}

	})

}

func (ts *AdvertisementRepoTestSuite) Test_GetById() {
	ts.Run("광고 고유값으로 조회 - 성공", func() {
		// given
		expectedResult := testAdvertisement
		id := 1
		rows := ts.mock.NewRows([]string{"id", "name", "image_url", "landing_url", "weight", "target_country", "target_gender", "reward"})
		for _, ad := range expectedResult {
			rows = rows.AddRow(ad.ID, ad.Name, ad.ImageUrl, ad.LandingUrl, ad.Weight, ad.TargetCountry, ad.TargetGender, ad.Reward)
		}
		ts.mock.ExpectQuery("SELECT (.+) FROM advertisement WHERE id = \\?").
			WithArgs(id).
			WillReturnRows(rows)
		// when
		result, err := ts.advertisementRepo.GetById(context.Background(), id)
		if err != nil {
			ts.Fail("repository error - %v", err)
		}
		// then
		if err := ts.mock.ExpectationsWereMet(); err != nil {
			ts.Fail("expectations were not met", err)
		}
		ts.Equal(expectedResult[0], result)
	})
	ts.Run("광고 고유값으로 조회 - 실패", func() {
		// given
		id := 1
		// when
		ts.mock.ExpectQuery("SELECT (.+) FROM advertisement WHERE id = \\?").
			WithArgs(id).
			WillReturnError(domain.ErrNotFound)
		_, err := ts.advertisementRepo.GetById(context.Background(), id)
		// then
		if err == nil {
			ts.Fail("expected prepare error, got", err)
		}
		if err := ts.mock.ExpectationsWereMet(); err != nil {
			ts.Fail("expectations were not met", err)
		}
	})
}
