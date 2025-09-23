package models

import (
	"database/sql"
	"time"
)

type AccountInfo struct {
	ID            uint64 `json:"uid" gorm:"column:uid"` //默认为主键
	UserName      string `json:"username" gorm:"column:username;index"`
	RealName      string `json:"realname" gorm:"column:realname;index"`
	Email         string
	PermGroupID   uint32       //用户类型
	Password      string       `json:"-"`
	StudentID     string       `json:"studentID" gorm:"index"`     //学号
	Major         string       `json:"major"`         //专业
	Department    string       `json:"department"`    //院系
	Grade         string       `json:"grade"`         //年级
	PhoneNumber   string       `json:"phoneNumber"`   //手机号
	Avatar        string       `json:"avatarFile"`    //头像文件名
	WechatOpenID  string       `json:"wechatOpenID"`  //微信openid，留给第三方做的，可以不是微信
	TwoFactorAuth string       `json:"twoFactorAuth"` //双因素认证密钥  F:7J64V3P3E77J3LKNUGSZ5QANTLRLTKVL
	CreatedAt     time.Time    `json:"createdAt" gorm:"index"`
	UpdatedAt     time.Time    `json:"updatedAt"`
	DeletedAt     sql.NullTime `json:"deletedAt,omitempty"`
}

