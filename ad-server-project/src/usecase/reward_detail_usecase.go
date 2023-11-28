package usecase

import (
	"ad-server-project/src/domain"
	"ad-server-project/src/domain/model"
	"context"
	"database/sql"
	"fmt"
)

type rewardDetailUsecase struct {
	rewardDetailRepo  model.RewardDetailRepository
	userRepo          model.UserRepository
	transactionRepo   model.TransactionRepository
	advertisementRepo model.AdvertisementRepository
}

func NewRewardDetailUsecase(r model.RewardDetailRepository, u model.UserRepository, t model.TransactionRepository, a model.AdvertisementRepository) model.RewardDetailUsecase {
	return &rewardDetailUsecase{
		rewardDetailRepo:  r,
		userRepo:          u,
		transactionRepo:   t,
		advertisementRepo: a,
	}
}

func (r *rewardDetailUsecase) EarnRewardDetail(c context.Context, adId int, reward int, userId int) error {
	rewardType := model.Plus

	user, err := r.userRepo.GetById(c, userId)
	if err != nil {
		return err
	}
	// 트랜잭션
	err = r.transactionRepo.Transaction(context.Background(), func(tx *sql.Tx) error {
		// 광고에 해당하는 리워드 값이 맞는지 확인
		ad, err := r.advertisementRepo.GetById(c, adId)
		if err != nil {
			return err
		}
		if ad.Reward != reward {
			return domain.ErrBadParamInput
		}
		// 유저의 리워드 정보 업데이트
		err = r.rewardDetailRepo.EarnRewardDetail(c, adId, reward, userId, rewardType)
		if err != nil {
			return err
		}
		user.Reward += reward

		err = r.userRepo.UpdateReward(c, user)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (r *rewardDetailUsecase) DeductRewardDetail(c context.Context, reward int, userId int) error {
	rewardType := model.Minus

	user, err := r.userRepo.GetById(c, userId)
	if err != nil {
		return err
	}
	// 트랜잭션
	err = r.transactionRepo.Transaction(context.Background(), func(tx *sql.Tx) error {
		err = r.rewardDetailRepo.DeductRewardDetail(c, reward, userId, rewardType)
		if err != nil {
			return err
		}

		if user.Reward < reward {
			return fmt.Errorf("no more rewards can be deducted")
		}
		user.Reward -= reward

		err = r.userRepo.UpdateReward(c, user)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (r *rewardDetailUsecase) GetRecent(c context.Context, userId int) ([]model.RewardDetail, error) {
	return r.rewardDetailRepo.GetRecent(c, userId)
}

func (r *rewardDetailUsecase) GetRewardBalance(c context.Context, userId int) (int, error) {
	user, err := r.userRepo.GetById(c, userId)
	if err != nil {
		return -1, err
	}
	return user.Reward, nil
}
