package usecase

import (
	"ad-server-project/src/domain/model"
	"ad-server-project/src/domain/model/mocks"
	"context"
	"fmt"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AdvertisementUsecaseTestSuite struct {
	suite.Suite
	AdvertisementUsecase  *advertisementUsecase
	mockAdvertisementRepo *mocks.AdvertisementRepository
}

func (ts *AdvertisementUsecaseTestSuite) SetupTest() {
	ts.mockAdvertisementRepo = new(mocks.AdvertisementRepository)
	ts.AdvertisementUsecase = &advertisementUsecase{
		advertisementRepo: ts.mockAdvertisementRepo,
	}
}

func TestAdvertisementUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(AdvertisementUsecaseTestSuite))
}

func generateTestAd(N int) []model.Advertisement {
	var ads []model.Advertisement

	for i := 1; i <= N; i++ {
		ad := model.Advertisement{
			ID:            i,
			Name:          fmt.Sprintf("Ad%d", i),
			ImageUrl:      fmt.Sprint("http://example.com/image.jpg"),
			LandingUrl:    fmt.Sprint("http://example.com/landing"),
			Weight:        i,
			TargetCountry: "KR",
			TargetGender:  "W",
			Reward:        i,
		}
		ads = append(ads, ad)
	}

	return ads
}

func (ts *AdvertisementUsecaseTestSuite) Test_GetByCountryAndGender() {
	ts.Run("나라와 성별 정보에 맞는 광고 조회", func() {
		// given
		userId := 1
		ads := generateTestAd(5)
		ts.mockAdvertisementRepo.On("GetByCountryAndGender", mock.Anything, mock.Anything).Return(ads, nil)
		// when
		result, err := ts.AdvertisementUsecase.GetByCountryAndGender(context.Background(), userId, "W", "KR")
		ts.NoError(err)
		ts.Equal(3, len(result))
	})
}
