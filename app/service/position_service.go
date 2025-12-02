package service

import (
	"bit-labs.cn/flex-admin/app/model"
	"bit-labs.cn/flex-admin/app/repository"
	"bit-labs.cn/owl/provider/db"
	"bit-labs.cn/owl/provider/redis"
	"bit-labs.cn/owl/provider/router"
	"github.com/jinzhu/copier"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type CreatePositionReq struct {
	Name   string `json:"name" validate:"omitempty,min=2,max=32"`
	Remark string `json:"remark" validate:"omitempty,max=255"`
	Status int    `json:"status" validate:"omitempty,oneof=0 1"`
}

type UpdatePositionReq struct {
	ID uint `json:"id,string"`
	CreatePositionReq
}

type RetrievePositionReq struct {
	router.PageReq
	NameLike string `json:"nameLike" validate:"omitempty,max=32"`
	Status   int    `json:"status" validate:"omitempty,oneof=0 1"`
}

type PositionService struct {
	db.BaseRepository[model.Position]
	repo   repository.PositionRepositoryInterface
	locker redis.LockerFactory
}

func NewPositionService(repo repository.PositionRepositoryInterface, tx *gorm.DB, locker redis.LockerFactory) *PositionService {
	return &PositionService{BaseRepository: db.NewBaseRepository[model.Position](tx), repo: repo, locker: locker}
}

func (i *PositionService) CreatePosition(req *CreatePositionReq) error {
	l := i.locker.New()
	if err := l.Lock("position:create"); err != nil {
		return err
	}
	defer l.Unlock()
	var m model.Position
	_ = copier.Copy(&m, req)
	m.Status = 1
	return i.repo.Save(&m)
}

func (i *PositionService) UpdatePosition(req *UpdatePositionReq) error {
	l := i.locker.New()
	if err := l.Lock("position:update:" + cast.ToString(req.ID)); err != nil {
		return err
	}
	defer l.Unlock()
	m, err := i.repo.Detail(req.ID)
	if err != nil {
		return err
	}
	if err = copier.Copy(m, req); err != nil {
		return err
	}
	return i.repo.Save(m)
}

func (i *PositionService) DeletePosition(id uint) error {
	l := i.locker.New()
	if err := l.Lock("position:delete:" + cast.ToString(id)); err != nil {
		return err
	}
	defer l.Unlock()
	return i.BaseRepository.Delete(id)
}

func (i *PositionService) ChangeStatus(req *db.ChangeStatus) error {
	l := i.locker.New()
	if err := l.Lock("position:status:" + cast.ToString(req.ID)); err != nil {
		return err
	}
	defer l.Unlock()
	return i.BaseRepository.ChangeStatus(req)
}

func (i *PositionService) RetrievePositions(req *RetrievePositionReq) (count int64, list []model.Position, err error) {
	return i.repo.Retrieve(req.Page, req.PageSize, func(tx *gorm.DB) { db.AppendWhereFromStruct(tx, req) })
}

func (i *PositionService) Options() (list []repository.PositionItem, err error) {
	return i.repo.Options()
}
