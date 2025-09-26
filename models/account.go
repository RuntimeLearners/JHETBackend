package models

import (
	"database/sql"
	"time"
)

// 数据库中用户信息模型
type AccountInfo struct {

	// 用户账号信息

	ID           uint64 `json:"id" gorm:"column:id;primaryKey"`        // 默认为主键
	Email        string `json:"email" gorm:"column:email;index"`       // 邮箱
	PasswordHash string `json:"-" gorm:"column:password_hash"`         // 密码的哈希值
	UserName     string `json:"username" gorm:"column:username;index"` // 用户显示名称
	PermGroupID  uint32 `json:"permGroupID" gorm:"index"`              // 用户所在权限组
	AvatarFile   string `json:"avatarFile"`                            //头像文件名
	Activation   string `json:"-"`                                     //账户激活状态(保留,用于验证邮箱是否存在)

	// 用户关联信息

	RealName    string       `json:"realname" gorm:"column:realname;index"`
	StudentID   string       `json:"studentID" gorm:"index"` //学号/人员编号
	Major       string       `json:"major"`                  //专业
	Department  string       `json:"department"`             //部门/院系 学生和管理员均有此项
	Grade       string       `json:"grade"`                  //年级 F:2025
	PhoneNumber string       `json:"phoneNumber"`            //手机号
	CreatedAt   time.Time    `json:"createdAt" gorm:"index"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	DeletedAt   sql.NullTime `json:"deletedAt,omitempty"`

	// 备用

	TwoFactorAuth string `json:"twoFactorAuth"` //双因素认证密钥  F:7J64V3P3E77J3LKNUGSZ5QANTLRLTKVL
	WechatOpenID  string `json:"wechatOpenID"`  //微信openid，留给第三方做的，可以不是微信
}
