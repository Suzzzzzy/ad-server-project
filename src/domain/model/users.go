package model

import (
	"context"
	"time"
)

type User struct {
	// 유저의 고유 아이디
	ID int `json:"id"`
	// 유저의 성별 정보 (M: 남자, F: 여자)
	Gender string `json:"gender"`
	// 유저의 국가 정보
	Country string `json:"country"`
	// 유저가 가진 reward 값
	Reward int `json:"reward"`
	CreatedAt time.Time `json:"created_at"`
}

type UserRepository interface {
	UpdateReward(c context.Context, user User) error
	GetById(c context.Context, userId int) (User, error)
}