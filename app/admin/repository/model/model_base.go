package model

import (
	"gorm.io/gorm"
	"time"
)

type ModelBase struct {
	ID        uint           `gorm:"primarykey" json:"id,string"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
