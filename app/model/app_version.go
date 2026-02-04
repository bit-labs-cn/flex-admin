package model

import "time"

type AppVersion struct {
	ID          int64      `gorm:"primaryKey;column:id" json:"id,string"`
	Version     *string    `gorm:"column:version" json:"version,omitempty"`
	VersionName *string    `gorm:"column:version_name" json:"versionName,omitempty"`
	ApkURL      *string    `gorm:"column:apk_url" json:"apkUrl,omitempty"`
	ApkType     *int32     `gorm:"column:apk_type" json:"apkType,omitempty"`
	AuditState  *int32     `gorm:"column:audit_state" json:"auditState,omitempty"`
	Content     *string    `gorm:"column:content" json:"content,omitempty"`
	Remark      *string    `gorm:"column:remark" json:"remark,omitempty"`
	Status      *int32     `gorm:"column:status" json:"status,omitempty"`
	CreateTime  *time.Time `gorm:"column:create_time" json:"createTime,omitempty"`
	UpdateTime  *time.Time `gorm:"column:update_time" json:"updateTime,omitempty"`
	CreatorID   *int64     `gorm:"column:creator_id" json:"creatorId,omitempty"`
	ModifierID  *int64     `gorm:"column:modifier_id" json:"modifierId,omitempty"`
	AutiTime    *time.Time `gorm:"column:auti_time" json:"autiTime,omitempty"`
	AutiID      *int64     `gorm:"column:auti_id" json:"autiId,omitempty"`
	UseOrgID    *int64     `gorm:"column:use_org_id" json:"useOrgId,omitempty"`
	CreateOrgID *int64     `gorm:"column:create_org_id" json:"createOrgId,omitempty"`
	ExtField1   *string    `gorm:"column:ext_field1" json:"extField1,omitempty"`
	ExtField2   *string    `gorm:"column:ext_field2" json:"extField2,omitempty"`
	ExtField3   *string    `gorm:"column:ext_field3" json:"extField3,omitempty"`
	ExtField4   *string    `gorm:"column:ext_field4" json:"extField4,omitempty"`
	ExtField5   *string    `gorm:"column:ext_field5" json:"extField5,omitempty"`
}

func (AppVersion) TableName() string {
	return "app_version"
}
