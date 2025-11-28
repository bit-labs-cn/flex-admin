package model

type OperationLog struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	UserId    int    `json:"userId" gorm:"column:user_id"`
	UserName  string `json:"userName" gorm:"column:user_name"`
	UserType  string `json:"userType" gorm:"column:user_type"`
	Method    string `json:"method" gorm:"column:method"`
	Path      string `json:"path" gorm:"column:path"`
	Status    int    `json:"status" gorm:"column:status"`
	CostMs    int    `json:"costMs" gorm:"column:cost_ms"`
	Ip        string `json:"ip" gorm:"column:ip"`
	UserAgent string `json:"userAgent" gorm:"column:user_agent"`
	ReqBody   string `json:"reqBody" gorm:"column:req_body;type:text"`
	CreatedAt int    `json:"createdAt" gorm:"column:created_at"`
}

func (i OperationLog) TableName() string {
	return "admin_operation_log"
}
