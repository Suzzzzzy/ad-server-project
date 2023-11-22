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

func (a advertisementUsecase) GetByCountryAndGender(c context.Context, user *model.User) ([]model.Advertisement, error) {
	res, err := a.advertisementRepo.GetByCountryAndGender(c, user)
	if err != nil {
		println("advertisementUsecase > advertisementRepo.GetByCountryAndGender Error \n", err)
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil
	}

	adWithWeightList := model.ConvertAdwithWeight(res)

	result := pickAdRandWithWeight(adWithWeightList, 3)

	return result, nil
}

func pickAdRandWithWeight(list []model.AdWithWeight, num int) []model.Advertisement {
	var result []model.Advertisement

	totalWeight := 0
	for _, item := range list {
		totalWeight += item.Weight
	}

	for i := 0; i < num; i++ {
		// 랜덤하게 선택된 확률
		randomProb := rand.Float64()
		// 누적 확률 변수
		cumulativeProb := 0.0

		for _, item := range list {
			cumulativeProb += float64(item.Weight) / float64(totalWeight)
			if randomProb <= cumulativeProb {
				result = append(result, item.Ad)
				break
			}
		}
	}

	return result
}
