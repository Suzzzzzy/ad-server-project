package model

type User struct {
	// 유저의 고유 아이디
	ID int `json:"id" gorm:"primary_key"`
	// 유저의 성별 정보 (M: 남자, F: 여자)
	Gender string `json:"gender"`
	// 유저의 국가 정보
	Country string `json:"country"`
}
