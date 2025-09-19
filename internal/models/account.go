package models

type AccountInfo struct {
	UID      uint64 `json:"uid"`
	UserName string `json:"username"`
	Email    string
	Perms    uint64
}
