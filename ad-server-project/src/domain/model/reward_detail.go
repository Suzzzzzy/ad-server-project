package model

import (
	"context"
	"github.com/guregu/null"
	"time"
)

// RewardDetail 리워드 상세 내역
type RewardDetail struct {
	ID int `json:"id"`
	// 리워드를 적립한 광고 고유 아이디
	AdId null.Int `json:"ad_id"`
	// 유저 아이디
	UserId int `json:"user_id"`
	// 적립 및 차감된 리워드 값
	Reward int `json:"reward"`
	// 적립/차감 타입(plus/minus)
	RewardType string `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

type RewardDetailRepository interface {
	// EarnRewardDetail 유저의 리워드 적립
	EarnRewardDetail(c context.Context, adId int, reward int, userId int, rewardType string) error
	// DeductRewardDetail 유저의 리워드 차감
	DeductRewardDetail(c context.Context, reward int, userId int, rewardType string) error
	// GetRecent 최근 일주일 리워드 내역 조회
	GetRecent(c context.Context, userId int) ([]RewardDetail, error)
}

type RewardDetailUsecase interface {
	EarnRewardDetail(c context.Context, adId int, reward int, userId int) error
	DeductRewardDetail(c context.Context, reward int, userId int) error
	GetRecent(c context.Context, userId int) ([]RewardDetail, error)
}
