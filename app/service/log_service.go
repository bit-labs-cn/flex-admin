package service

import (
	"bit-labs.cn/flex-admin/app/model"
	"bit-labs.cn/flex-admin/app/repository"
	"bit-labs.cn/owl/provider/router"
	validatorv10 "github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type LogService struct {
	logRepo  repository.LogRepositoryInterface // 日志仓储接口
	validate *validatorv10.Validate
}

func NewLogService(repo repository.LogRepositoryInterface, validate *validatorv10.Validate) *LogService {
	return &LogService{logRepo: repo, validate: validate}
}

// 记录登录日志

type CreateOperationReq struct {
	UserId    int    `json:"userId" validate:"required,gte=1"`           // 用户编号
	UserName  string `json:"userName" validate:"required,max=64"`        // 用户名称
	UserType  string `json:"userType" validate:"required,max=32"`        // 用户类型（user/super_admin）
	Method    string `json:"method" validate:"required,max=16"`          // 请求方法
	Path      string `json:"path" validate:"required,max=255"`           // 请求路径
	ApiName   string `json:"apiName" validate:"omitempty,max=64"`        // 接口中文名称
	Status    int    `json:"status" validate:"required,gte=100,lte=599"` // 响应状态码（HTTP）
	CostMs    int    `json:"costMs" validate:"omitempty,gte=0"`          // 耗时毫秒
	Ip        string `json:"ip" validate:"omitempty,max=64"`             // 客户端 IP
	UserAgent string `json:"userAgent" validate:"omitempty,max=255"`     // 客户端 UA
	ReqBody   string `json:"reqBody" validate:"omitempty"`               // 请求体（文本）
}

type CreateLoginLogReq struct {
	UserId    int    `json:"userId" validate:"required,gte=1"`       // 用户编号
	UserName  string `json:"userName" validate:"required,max=64"`    // 用户名称
	UserType  string `json:"userType" validate:"required,max=32"`    // 用户类型（user/super_admin）
	LoginTime int    `json:"loginTime" validate:"required,gte=1"`    // 登录时间（Unix 秒）
	Ip        string `json:"ip" validate:"omitempty,max=64"`         // 客户端 IP
	UserAgent string `json:"userAgent" validate:"omitempty,max=255"` // 客户端 UA
}

func (i *LogService) RecordLogin(req *CreateLoginLogReq) error {
	if err := i.validate.Struct(req); err != nil {
		return err
	}

	var log model.LoginLog
	err := copier.Copy(&log, req)
	if err != nil {
		return err
	}

	return i.logRepo.SaveLoginLog(&log)
}

func (i *LogService) RecordOperation(req *CreateOperationReq) error {
	if err := i.validate.Struct(req); err != nil {
		return err
	}

	var log model.OperationLog
	err := copier.Copy(&log, req)
	if err != nil {
		return err
	}

	return i.logRepo.SaveOperationLog(&log)
}

type RetrieveLoginLogsReq struct {
	router.PageReq
	UserName string `json:"userName"` // 用户名
	Ip       string `json:"ip"`       // IP 地址
	UserType string `json:"userType"` // 用户类型（user/super_admin）
	Start    int    `json:"start"`    // 开始时间（Unix 秒）
	End      int    `json:"end"`      // 结束时间（Unix 秒）
}

func (i *LogService) RetrieveLoginLogs(req *RetrieveLoginLogsReq) (count int64, list []model.LoginLog, err error) {
	if err := i.validate.Struct(req); err != nil {
		return 0, nil, err
	}

	return i.logRepo.RetrieveLoginLogs(req.Page, req.PageSize, func(tx *gorm.DB) {
		if req.UserName != "" {
			tx.Where("user_name = ?", req.UserName)
		}
		if req.Ip != "" {
			tx.Where("ip = ?", req.Ip)
		}
		if req.UserType != "" {
			tx.Where("user_type = ?", req.UserType)
		}
		if req.Start > 0 && req.End > 0 {
			tx.Where("login_time BETWEEN ? AND ?", req.Start, req.End)
		}
	})
}

type RetrieveOperationLogsReq struct {
	router.PageReq
	UserName string `json:"userName"` // 用户名
	Path     string `json:"path"`     // 请求路径（模糊）
	Method   string `json:"method"`   // 请求方法
	Status   *int   `json:"status"`   // 状态码
	Start    int    `json:"start"`    // 开始时间（Unix 秒）
	End      int    `json:"end"`      // 结束时间（Unix 秒）
}

func (i *LogService) RetrieveOperationLogs(req *RetrieveOperationLogsReq) (count int64, list []model.OperationLog, err error) {
	if err := i.validate.Struct(req); err != nil {
		return 0, nil, err
	}
	return i.logRepo.RetrieveOperationLogs(req.Page, req.PageSize, func(tx *gorm.DB) {
		if req.UserName != "" {
			tx.Where("user_name = ?", req.UserName)
		}
		if req.Path != "" {
			tx.Where("path LIKE ?", "%"+req.Path+"%")
		}
		if req.Method != "" {
			tx.Where("method = ?", req.Method)
		}
		if req.Status != nil {
			tx.Where("status = ?", *req.Status)
		}
		if req.Start > 0 && req.End > 0 {
			tx.Where("created_at BETWEEN ? AND ?", req.Start, req.End)
		}
	})
}
