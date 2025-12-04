package service

import (
	"time"

	"bit-labs.cn/flex-admin/app/model"
	"bit-labs.cn/owl/provider/router"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LogService struct {
	db *gorm.DB
}

func NewLogService(tx *gorm.DB) *LogService {
	return &LogService{db: tx}
}

func (i *LogService) RecordLogin(ctx *gin.Context, user *model.User) error {
	ip := ctx.ClientIP()
	ua := ctx.GetHeader("User-Agent")
	t := int(time.Now().Unix())
	uType := "user"
	if user != nil && user.IsSuperAdmin {
		uType = "super_admin"
	}
	uId := 0
	uName := ""
	if user != nil {
		uId = int(user.ID)
		uName = user.Username
	}
	log := &model.LoginLog{
		Ip:        ip,
		LoginTime: t,
		UserId:    uId,
		UserName:  uName,
		UserType:  uType,
		UserAgent: ua,
	}
	return i.db.Create(log).Error
}

type CreateOperationReq struct {
	UserId    int    `json:"userId"`
	UserName  string `json:"userName"`
	UserType  string `json:"userType"`
	Method    string `json:"method"`
	Path      string `json:"path"`
	ApiName   string `json:"apiName"`
	Status    int    `json:"status"`
	CostMs    int    `json:"costMs"`
	Ip        string `json:"ip"`
	UserAgent string `json:"userAgent"`
	ReqBody   string `json:"reqBody"`
}

func (i *LogService) RecordOperation(req *CreateOperationReq) error {

	log := &model.OperationLog{
		UserId:    req.UserId,
		UserName:  req.UserName,
		UserType:  req.UserType,
		Method:    req.Method,
		Path:      req.Path,
		ApiName:   req.ApiName,
		Status:    req.Status,
		CostMs:    req.CostMs,
		Ip:        req.Ip,
		UserAgent: req.UserAgent,
		ReqBody:   req.ReqBody,
		CreatedAt: int(time.Now().Unix()),
	}
	return i.db.Create(log).Error
}

type RetrieveLoginLogsReq struct {
	router.PageReq
	UserName string `json:"userName"`
	Ip       string `json:"ip"`
	UserType string `json:"userType"`
	Start    int    `json:"start"`
	End      int    `json:"end"`
}

func (i *LogService) RetrieveLoginLogs(req *RetrieveLoginLogsReq) (count int64, list []model.LoginLog, err error) {
	tx := i.db.Model(&model.LoginLog{})
	if req.UserName != "" {
		tx = tx.Where("user_name = ?", req.UserName)
	}
	if req.Ip != "" {
		tx = tx.Where("ip = ?", req.Ip)
	}
	if req.UserType != "" {
		tx = tx.Where("user_type = ?", req.UserType)
	}
	if req.Start > 0 && req.End > 0 {
		tx = tx.Where("login_time BETWEEN ? AND ?", req.Start, req.End)
	}
	err = tx.Count(&count).Error
	if err != nil {
		return
	}
	err = tx.Order("id DESC").Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize).Find(&list).Error
	return
}

type RetrieveOperationLogsReq struct {
	router.PageReq
	UserName string `json:"userName"`
	Path     string `json:"path"`
	Method   string `json:"method"`
	Status   *int   `json:"status"`
	Start    int    `json:"start"`
	End      int    `json:"end"`
}

func (i *LogService) RetrieveOperationLogs(req *RetrieveOperationLogsReq) (count int64, list []model.OperationLog, err error) {
	tx := i.db.Model(&model.OperationLog{})
	if req.UserName != "" {
		tx = tx.Where("user_name = ?", req.UserName)
	}
	if req.Path != "" {
		tx = tx.Where("path LIKE ?", "%"+req.Path+"%")
	}
	if req.Method != "" {
		tx = tx.Where("method = ?", req.Method)
	}
	if req.Status != nil {
		tx = tx.Where("status = ?", *req.Status)
	}
	if req.Start > 0 && req.End > 0 {
		tx = tx.Where("created_at BETWEEN ? AND ?", req.Start, req.End)
	}
	err = tx.Count(&count).Error
	if err != nil {
		return
	}
	err = tx.Order("id DESC").Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize).Find(&list).Error
	return
}
