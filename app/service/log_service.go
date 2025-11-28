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

func (i *LogService) RecordOperation(ctx *gin.Context, user *model.User, status int, costMs int, reqBody string) error {
	ip := ctx.ClientIP()
	ua := ctx.GetHeader("User-Agent")
	path := ctx.FullPath()
	if path == "" {
		path = ctx.Request.URL.Path
	}
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
	log := &model.OperationLog{
		UserId:    uId,
		UserName:  uName,
		UserType:  uType,
		Method:    ctx.Request.Method,
		Path:      path,
		Status:    status,
		CostMs:    costMs,
		Ip:        ip,
		UserAgent: ua,
		ReqBody:   reqBody,
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
