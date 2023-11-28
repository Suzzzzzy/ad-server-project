package repository

import (
	"ad-server-project/src/domain/model"
	"database/sql"
	"fmt"
	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
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

func TestAdvertisementRepoTestsute(t *testing.T) {
	suite.Run(t, new(AdvertisementRepoTestSuite))
}

func (ts *AdvertisementRepoTestSuite) Test_GetByCountryAndGender() {
	ts.Run("국가, 성별로 광고 조회", func() {

	})
}
