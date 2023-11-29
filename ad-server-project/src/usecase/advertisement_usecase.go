package usecase

import (
	"ad-server-project/src/adapter"
	"ad-server-project/src/domain/model"
	"context"
	"math/rand"
	"sort"
	"time"
)

type advertisementUsecase struct {
	advertisementRepo model.AdvertisementRepository
}

func NewAdvertisementUsecase(a model.AdvertisementRepository) model.AdvertisementUsecase {
	return &advertisementUsecase{
		advertisementRepo: a,
	}
}

func (a *advertisementUsecase) GetByCountryAndGender(c context.Context, userId int, userGender string, userCountry string) ([]model.AdvertisementMinInfo, error) {
	user := &model.User{
		ID:      userId,
		Gender:  userGender,
		Country: userCountry,
	}
	res, err := a.advertisementRepo.GetByCountryAndGender(c, user)
	if err != nil {
		println("advertisementUsecase > advertisementRepo.GetByCountryAndGender Error \n", err)
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil
	}

	var pickedAd []model.Advertisement
	num := 3
	if num > len(res) {
		num = len(res)
	}

	// userId로 정책 선택
	methodChoice := user.ID % 4
	switch methodChoice {
	case 0:
		pickedAd = pickAdByRandom(res, num)
	case 1:
		pickedAd = pickAdByWeight(res, num)
	case 2:
		pickedAd = pickAdByPctr(res, user.ID, num)
	case 3:
		pickedAd = pickAdByWeightPctrMixed(res, user.ID, num)
	}

	result := model.ConvertAdMinInfo(pickedAd)

	return result, nil
}

// pickAdRand random 정책: 랜덤으로 정렬하는 정책
func pickAdByRandom(list []model.Advertisement, num int) []model.Advertisement {
	listLength := len(list)

	rand.Seed(time.Now().UnixNano())

	result := make([]model.Advertisement, num)

	selectedIndexes := make(map[int]bool)

	for i := 0; i < num; {
		randomIndex := rand.Intn(listLength)
		// 이미 선택된 인덱스인지 확인
		if !selectedIndexes[randomIndex] {
			selectedIndexes[randomIndex] = true

			result[i] = list[randomIndex]
			i++
		}
	}
	return result
}

// pickAdByWeight weight 정책: weight 기반의 정책
func pickAdByWeight(list []model.Advertisement, num int) []model.Advertisement {
	adWithWeightList := model.ConvertAdwithWeight(list)

	var result []model.Advertisement

	totalWeight := 0
	for _, item := range adWithWeightList {
		totalWeight += item.Weight
	}

	// Precalculate cumulative probabilities
	cumulativeProbs := make([]float64, len(adWithWeightList))
	cumProb := 0.0
	for i, item := range adWithWeightList {
		cumProb += float64(item.Weight) / float64(totalWeight)
		cumulativeProbs[i] = cumProb
	}

	for i := 0; i < num; i++ {
		// Random probability
		randomProb := rand.Float64()

		// Find the index where randomProb falls in the cumulative probabilities
		index := sort.Search(len(cumulativeProbs), func(j int) bool {
			return cumulativeProbs[j] >= randomProb
		})

		// Append the selected advertisement to the result
		result = append(result, adWithWeightList[index].Ad)
	}

	return result

}

// pickAdByPctr pctr 정책: 예측된 CTR의 내림차순으로 정렬하는 정책
func pickAdByPctr(list []model.Advertisement, userId int, num int) []model.Advertisement {
	adIdList := make([]int, len(list))
	for _, ad := range list {
		adIdList = append(adIdList, ad.ID)
	}

	pctr, err := adapter.SendCtrPredictionServer(userId, adIdList)
	if err != nil {
		return nil
	}

	// 광고의 Id와 pctr 맵핑
	adIdToPctr := make(map[int]float64)
	for i, ad := range adIdList {
		adIdToPctr[ad] = pctr.PCTR[i]
	}

	// 광고 정렬
	sort.Slice(list, func(i, j int) bool {
		return adIdToPctr[list[i].ID] > adIdToPctr[list[j].ID]
	})

	return list[:num]
}

// pickAdByWeightPctrMixed weight_pctr_mixed 정책: 예측된 CTR이 가장 높은 광고를 첫 번째에 위치하고 나머지 두 광고는 weight 기반으로 정렬하는 정책
func pickAdByWeightPctrMixed(list []model.Advertisement, userId int, num int) []model.Advertisement {
	var result []model.Advertisement
	// CTR 높은 광고 1개
	adByPctr := pickAdByPctr(list, userId, 1)
	result = append(result, adByPctr...)

	// 나머지는 weight 기반 광고
	adByWeight := pickAdByWeight(list, num-1)
	result = append(result, adByWeight...)

	return result
}

func (a *advertisementUsecase) UpdateReward(c context.Context, id int, reward int) error {
	ad, err := a.advertisementRepo.GetById(c, id)
	if err != nil {
		return err
	}
	return a.advertisementRepo.UpdateReward(c, ad.ID, reward)
}
