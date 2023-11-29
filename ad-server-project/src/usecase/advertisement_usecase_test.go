package usecase

import (
	"ad-server-project/src/domain"
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
	advertisementUsecase  *advertisementUsecase
	mockAdvertisementRepo *mocks.AdvertisementRepository
}

func (ts *AdvertisementUsecaseTestSuite) SetupTest() {
	ts.mockAdvertisementRepo = new(mocks.AdvertisementRepository)
	ts.advertisementUsecase = &advertisementUsecase{
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
		result, err := ts.advertisementUsecase.GetByCountryAndGender(context.Background(), userId, "W", "KR")
		ts.NoError(err)
		ts.Equal(3, len(result))
	})
}

func (ts *AdvertisementUsecaseTestSuite) Test_UpdateReward() {
	// given
	ad := generateTestAd(1)[0]
	ts.Run("광고의 리워드값 업데이트 - 성공", func() {
		ts.SetupTest()
		ts.mockAdvertisementRepo.On("GetById", mock.Anything, mock.Anything).Return(ad, nil)
		ts.mockAdvertisementRepo.On("UpdateReward", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		// when
		result := ts.advertisementUsecase.UpdateReward(context.Background(), ad.ID, 10)
		// then
		ts.NoError(result)
	})
	ts.Run("광고의 리워드값 업데이트 - 실패(해당 광고 존재하지 않음)", func() {
		ts.SetupTest()
		ts.mockAdvertisementRepo.On("GetById", mock.Anything, mock.Anything).Return(model.Advertisement{}, domain.ErrNotFound)
		// when
		result := ts.advertisementUsecase.UpdateReward(context.Background(), ad.ID, 10)
		// then
		ts.Error(result)
		ts.Equal(domain.ErrNotFound, result)
	})
	ts.Run("광고의 리워드값 업데이트 - 실패(리워드값이 동일하여 수정되지 않음)", func() {
		ts.SetupTest()
		ts.mockAdvertisementRepo.On("GetById", mock.Anything, mock.Anything).Return(ad, nil)
		ts.mockAdvertisementRepo.On("UpdateReward", mock.Anything, mock.Anything, mock.Anything).Return(fmt.Errorf("No change in reward value"))
		// when
		result := ts.advertisementUsecase.UpdateReward(context.Background(), ad.ID, 10)
		// then
		ts.Error(result)
		ts.Equal("No change in reward value", result.Error())
	})
}
