package model

import (
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        uint           `gorm:"<-:create" gorm:"primarykey" json:"id,string,omitempty"`
	CreatedAt time.Time      `gorm:"<-:create" json:"createdAt,omitempty"`
	UpdatedAt time.Time      `json:"updatedAt" json:"updatedAt,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}
