package models

import "time"

type AccountInfo struct {
	UID          uint64 `json:"uid"`
	UserName     string `json:"username"`
	RealName     string `json:"realname"`
	Email        string
	PermGroupID  uint32    //用户类型？
	Password     string    `json:"-"`
	StudentID    string    `json:"studentID"`    //学号
	Major        string    `json:"major"`        //专业
	Department   string    `json:"department"`   //院系
	Grade        string    `json:"grade"`        //年级
	PhoneNumber  string    `json:"phoneNumber"`  //手机号
	Avatar       string    `json:"avatarFile"`   //头像文件名
	WechatOpenID string    `json:"wechatOpenID"` //微信openid
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	DeletedAt    time.Time `json:"deletedAt,omitempty"`
}
