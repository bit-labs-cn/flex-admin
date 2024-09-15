package db

import (
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
	CreatorID string `gorm:"comment:创建人" json:"creatorID"` // 创建人
	UpdaterID string `gorm:"comment:修改人" json:"updaterID"` // 修改人
}

func (i *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	//idGenerator, err := util.NewSnowflake(1, 3)
	//if err != nil {
	//	return err
	//}
	//i.ID = cast.ToUint(idGenerator.NextVal())
	//i.ID = uint(rand.Int())
	return
}
