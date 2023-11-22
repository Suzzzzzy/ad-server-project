package model

import (
	"context"
)

type Advertisement struct {
	// 광고를 구분하는 고유 아이디
	ID int `json:"id" gorm:"primary_key"`
	// 광고의 이름
	Name string `json:"name"`
	// 유저에게 보일 광고의 이미지 주소
	ImageUrl string `json:"image_url"`
	// 광고를 클릭했을 때 최종으로 유저가 랜딩 되어야 할 광고주 페이지
	LandingUrl string `json:"landing_url"`
	// 광고의 송출 가중치
	Weight int `json:"weight"`
	// 광고가 송출 가능한 국가 정보
	TargetCountry string `json:"target_country"`
	// 광고의 성별 타게팅 정보 (M: 남자 타게팅, F: 여자 타게팅)
	TargetGender string `json:"target_gender"`
	// 광고를 클릭했을 때 받을 수 있는 리워드
	Reward int `json:"reward"`
}

type AdWithWeight struct {
	Ad     Advertisement
	Weight int
}

type AdvertisementMinInfo struct {
	ImageUrl   string `json:"image_url"`
	LandingUrl string `json:"landing_url"`
	Reward     int    `json:"reward"`
}

type AdvertisementRepository interface {
	GetByCountryAndGender(c context.Context, user *User) ([]Advertisement, error)
}

type AdvertisementUsecase interface {
	GetByCountryAndGender(c context.Context, user *User) ([]AdvertisementMinInfo, error)
}

func ConvertAdwithWeight(list []Advertisement) []AdWithWeight {
	var result []AdWithWeight

	for _, ad := range list {
		result = append(result, AdWithWeight{
			Ad:     ad,
			Weight: ad.Weight,
		})
	}
	return result
}

func ConvertAdMinInfo(list []Advertisement) []AdvertisementMinInfo {
	var result []AdvertisementMinInfo
	for _, ad := range list {
		result = append(result, AdvertisementMinInfo{
			ImageUrl:   ad.ImageUrl,
			LandingUrl: ad.LandingUrl,
			Reward:     ad.Reward,
		})
	}
	return result
}
