package usecase

import (
	"ad-server-project/src/adapter"
	"ad-server-project/src/domain/model"
	"ad-server-project/src/utils"
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

func (a advertisementUsecase) GetByCountryAndGender(c context.Context, user *model.User) ([]model.AdvertisementMinInfo, error) {
	res, err := a.advertisementRepo.GetByCountryAndGender(c, user)
	if err != nil {
		println("advertisementUsecase > advertisementRepo.GetByCountryAndGender Error \n", err)
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil
	}

	pickedAd := pickAdByWeight(res, 3)

	result := model.ConvertAdMinInfo(pickedAd)

	return result, nil
}

// pickAdRand random 정책: 랜덤으로 정렬하는 정책
func pickAdByRandom(list []model.Advertisement, num int) []model.Advertisement {
	listLength := len(list)
	if num > listLength {
		num = listLength
	}

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
	if num > len(list){
		num = len(list)
	}
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

// pickAdByPctr pctr 정책: 예측된 CTR의 내림차순으로 정렬하는 정책
func pickAdByPctr(list []model.Advertisement, userId int, num int) []model.Advertisement {
	if num > len(list){
		num = len(list)
	}

	var adIdList []int
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

	sortedAdId := utils.SortMapByValue(adIdToPctr)

	var result []model.Advertisement
	for i :=0; i < num; i++ {
		for _, ad := range list {
			if sortedAdId[i].Key == ad.ID {
				result = append(result, ad)
				break
			}
		}
	}

	return result
}

// pickAdByWeightPctrMixed weight_pctr_mixed 정책: 예측된 CTR이 가장 높은 광고를 첫 번째에 위치하고 나머지 두 광고는 weight 기반으로 정렬하는 정책
func pickAdByWeightPctrMixed(list []model.Advertisement, num int) []model.Advertisement {

}