package model

import (
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        uint           `gorm:"<-:create" gorm:"primarykey" json:"id,string"`
	CreatedAt time.Time      `gorm:"<-:create" json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
