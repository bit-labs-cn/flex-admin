package model

type Position struct {
	Base
	Name   string `json:"name"`
	Remark string `json:"remark"`
	Status int    `json:"status"`
}

func (Position) TableName() string {
	return "admin_position"
}
