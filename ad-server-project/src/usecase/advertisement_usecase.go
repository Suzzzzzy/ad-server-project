package usecase

import (
	"ad-server-project/src/domain/model"
	"context"
	"math/rand"
)

type advertisementUsecase struct {
	advertisementRepo model.AdvertisementRepository
}

func NewAdvertisementUsecase(a model.AdvertisementRepository) model.AdvertisementUsecase {
	return &advertisementUsecase{
		advertisementRepo: a,
	}
}

func (a advertisementUsecase) GetByCountryAndGender(c context.Context, user *model.User) ([]model.AdvertisementMinInfo, error) {
	res, err := a.advertisementRepo.GetByCountryAndGender(c, user)
	if err != nil {
		println("advertisementUsecase > advertisementRepo.GetByCountryAndGender Error \n", err)
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil
	}

	pickedAd := pickAdRandWithWeight(res, 3)

	result := model.ConvertAdMinInfo(pickedAd)

	return result, nil
}

func pickAdRandWithWeight(list []model.Advertisement, num int) []model.Advertisement {
	adWithWeightList := model.ConvertAdwithWeight(list)

	var result []model.Advertisement

	totalWeight := 0
	for _, item := range adWithWeightList {
		totalWeight += item.Weight
	}

	for i := 0; i < num; i++ {
		// 랜덤하게 선택된 확률
		randomProb := rand.Float64()
		// 누적 확률 변수
		cumulativeProb := 0.0

		for _, item := range adWithWeightList {
			cumulativeProb += float64(item.Weight) / float64(totalWeight)
			if randomProb <= cumulativeProb {
				result = append(result, item.Ad)
				break
			}
		}
	}

	return result
}

func pickAdRand(list []model.)