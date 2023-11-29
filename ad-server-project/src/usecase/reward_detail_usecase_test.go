package usecase

import (
	"ad-server-project/src/domain"
	"ad-server-project/src/domain/model"
	"ad-server-project/src/domain/model/mocks"
	"context"
	"fmt"
	"github.com/guregu/null"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RewardDetailUsecaseTestSuite struct {
	suite.Suite
	rewardDetailUsecase   *rewardDetailUsecase
	mockRewardDetailRepo  *mocks.RewardDetailRepository
	mockUserRepo          *mocks.UserRepository
	mockTransactionRepo   *mocks.TransactionRepository
	mockAdvertisementRepo *mocks.AdvertisementRepository
}

func (ts *RewardDetailUsecaseTestSuite) SetupTest() {
	ts.mockRewardDetailRepo = new(mocks.RewardDetailRepository)
	ts.mockUserRepo = new(mocks.UserRepository)
	ts.mockTransactionRepo = new(mocks.TransactionRepository)
	ts.mockAdvertisementRepo = new(mocks.AdvertisementRepository)
	ts.rewardDetailUsecase = &rewardDetailUsecase{
		rewardDetailRepo:  ts.mockRewardDetailRepo,
		userRepo:          ts.mockUserRepo,
		transactionRepo:   ts.mockTransactionRepo,
		advertisementRepo: ts.mockAdvertisementRepo,
	}
}

func TestRewardDetailUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(RewardDetailUsecaseTestSuite))
}

func generateTestRewardDetail(N int, userId int, rewardType string) []model.RewardDetail {
	var rewardDetails []model.RewardDetail

	for i := 1; i <= N; i++ {
		ad := model.RewardDetail{
			ID:         i,
			AdId:       null.IntFrom(int64(i)),
			UserId:     userId,
			Reward:     i,
			RewardType: rewardType,
		}
		rewardDetails = append(rewardDetails, ad)
	}
	return rewardDetails
}

func (ts *RewardDetailUsecaseTestSuite) Test_EarnRewardDetail() {
	ad := model.Advertisement{ID: 1, Reward: 5}
	user := model.User{ID: 1, Reward: 0}

	ts.Run("유저의 리워드 적립 - 성공", func() {
		ts.mockUserRepo.On("GetById", mock.Anything, mock.Anything).Return(user, nil)
		ts.mockAdvertisementRepo.On("GetById", mock.Anything, mock.Anything).Return(ad, nil)
		ts.mockRewardDetailRepo.On("EarnRewardDetail", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
		ts.mockUserRepo.On("UpdateReward", mock.Anything, mock.Anything).Return(nil)
		ts.mockTransactionRepo.On("Transaction", mock.Anything, mock.Anything).Return(nil)

		result := ts.rewardDetailUsecase.EarnRewardDetail(context.Background(), ad.ID, ad.Reward, user.ID)
		ts.NoError(result)
	})
	ts.Run("유저의 리워드 적립 - 실패(유저 존재 X)", func() {
		ts.mockUserRepo.On("GetById", mock.Anything, mock.Anything).Return(user, domain.ErrUserNotFound)

		result := ts.rewardDetailUsecase.EarnRewardDetail(context.Background(), ad.ID, ad.Reward, user.ID)
		ts.Equal(domain.ErrUserNotFound, result)
	})
	ts.Run("유저의 리워드 적립 - 실패(광고와 리워드 정보가 일치하지 않음)", func() {
		ts.mockUserRepo.On("GetById", mock.Anything, mock.Anything).Return(user, nil)
		ts.mockAdvertisementRepo.On("GetById", mock.Anything, mock.Anything).Return(nil, domain.ErrNotFound)
		ts.mockTransactionRepo.On("Transaction", mock.Anything, mock.Anything).Return(domain.ErrNotFound)

		result := ts.rewardDetailUsecase.EarnRewardDetail(context.Background(), ad.ID, ad.Reward, user.ID)
		ts.Equal(domain.ErrNotFound, result)
	})
}

func (ts *RewardDetailUsecaseTestSuite) Test_DeductRewardDetail() {
	ad := model.Advertisement{ID: 1, Reward: 5}
	user := model.User{ID: 1}

	ts.Run("유저의 리워드 차감 - 성공", func() {
		ts.mockUserRepo.On("GetById", mock.Anything, mock.Anything).Return(user, nil)
		ts.mockRewardDetailRepo.On("DeductRewardDetail", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
		ts.mockUserRepo.On("UpdateReward", mock.Anything, mock.Anything).Return(nil)
		ts.mockTransactionRepo.On("Transaction", mock.Anything, mock.Anything).Return(nil)

		result := ts.rewardDetailUsecase.DeductRewardDetail(context.Background(), ad.Reward, user.ID)
		ts.NoError(result)
	})
	ts.Run("유저의 리워드 차감 - 실패(리워드는 음수값이 될 수 없다)", func() {
		ts.mockUserRepo.On("GetById", mock.Anything, mock.Anything).Return(user, nil)
		ts.mockRewardDetailRepo.On("DeductRewardDetail", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
		ts.mockUserRepo.On("UpdateReward", mock.Anything, mock.Anything).Return(fmt.Errorf("no more rewards can be deducted"))
		ts.mockTransactionRepo.On("Transaction", mock.Anything, mock.Anything).Return(fmt.Errorf("no more rewards can be deducted"))

		result := ts.rewardDetailUsecase.DeductRewardDetail(context.Background(), ad.Reward, user.ID)
		ts.Error(result)
		ts.Equal("no more rewards can be deducted", result.Error())
	})
}

func (ts *RewardDetailUsecaseTestSuite) Test_GetBalance() {
	user := model.User{ID: 1, Reward: 10}
	ts.Run("유저의 리워드 잔액 조회", func() {
		ts.mockUserRepo.On("GetById", mock.Anything, mock.Anything).Return(user, nil)

		result, err := ts.rewardDetailUsecase.GetRewardBalance(context.Background(), user.ID)
		ts.NoError(err)
		ts.Equal(user.Reward, result)
	})
	ts.Run("유저의 리워드 잔액 조회 - 실패(존재하지 않는 유저)", func() {
		ts.mockUserRepo.On("GetById", mock.Anything, mock.Anything).Return(model.User{}, domain.ErrUserNotFound)

		_, err := ts.rewardDetailUsecase.GetRewardBalance(context.Background(), user.ID)
		ts.Error(err)
	})
}
