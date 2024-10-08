package model

type LoginLog struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Ip        string `json:"ip" gorm:"column:ip"`
	LoginTime int    `json:"login_time" gorm:"column:login_time"`
	UserId    int    `json:"user_id" gorm:"column:user_id"`
	UserName  string `json:"user_name" gorm:"column:user_name"`
	UserType  string `json:"user_type" gorm:"column:user_type"`
	UserAgent string `json:"user_agent" gorm:"column:user_agent"`
}
