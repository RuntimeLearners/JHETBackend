package models

import "github.com/bits-and-blooms/bitset"

type PermissionGroup struct {
	// 权限组名称
	Name           string
	PermissionData []byte
	Permissions    bitset.BitSet
}
