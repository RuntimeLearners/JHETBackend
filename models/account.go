package models

import (
	"database/sql"
	"time"
)

// 数据库中用户信息模型
type AccountInfo struct {
<<<<<<< HEAD

	// 用户账号信息

	ID           uint64 `json:"uid" gorm:"column:uid"`                       // 默认为主键
	Email        string `json:"email" gorm:"column:email;index"`             // 邮箱
	PasswordHash string `json:"-" gorm:"column:passwordHash"`                // 密码的哈希值
	UserName     string `json:"username" gorm:"column:username;index"`       // 用户显示名称
	PermGroupID  uint32 `json:"permGroupID" gorm:"column:permGroupID;index"` // 用户所在权限组
	AvatarFile   string `json:"avatarFile"`                                  //头像文件名

	// 用户关联信息

	RealName    string       `json:"realname" gorm:"column:realname;index"`
	StudentID   string       `json:"studentID" gorm:"index"` //学号
	Major       string       `json:"major"`                  //专业
	Department  string       `json:"department"`             //院系
	Grade       string       `json:"grade"`                  //年级
	PhoneNumber string       `json:"phoneNumber"`            //手机号
	CreatedAt   time.Time    `json:"createdAt" gorm:"index"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	DeletedAt   sql.NullTime `json:"deletedAt,omitempty"`

	// 备用

	TwoFactorAuth string `json:"twoFactorAuth"` //双因素认证密钥  F:7J64V3P3E77J3LKNUGSZ5QANTLRLTKVL
=======
	ID            uint64       `gorm:"column:id;primaryKey"` //默认为主键
	UserName      string       `json:"username" gorm:"column:username;index"`
	RealName      string       `json:"realname" gorm:"column:realname;index"`
	Email         string       `json:"email"`
	PermGroupID   uint32       `json:"permGroupID"`            //用户类型
	Password      string       `json:"-"`                      //密码的哈希值
	StudentID     string       `json:"studentID" gorm:"index"` //学号/人员编号
	Major         string       `json:"major"`                  //专业
	Department    string       `json:"department"`             //部门/院系 学生和管理员均有此项
	Grade         string       `json:"grade"`                  //年级 F:2025
	PhoneNumber   string       `json:"phoneNumber"`            //手机号
	Avatar        string       `json:"avatarFile"`             //头像文件名
	WechatOpenID  string       `json:"wechatOpenID"`           //微信openid，留给第三方做的，可以不是微信
	TwoFactorAuth string       `json:"twoFactorAuth"`          //双因素认证密钥  F:7J64V3P3E77J3LKNUGSZ5QANTLRLTKVL
	Activation    string       `json:"-"`                      //账户激活状态(保留,用于验证邮箱是否存在)
	CreatedAt     time.Time    `json:"createdAt" gorm:"index"`
	UpdatedAt     time.Time    `json:"updatedAt"`
	DeletedAt     sql.NullTime `json:"deletedAt,omitempty"`
>>>>>>> 1f4aefcfc69c5a7434c70f0b1a5ca47c73bd2f55
}
