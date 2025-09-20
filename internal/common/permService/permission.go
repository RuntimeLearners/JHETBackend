package permission

import "JHETBackend/internal/configs/database"

package permission

import (
	"database/sql"
	"fmt"
	"sync"
)

// 1. 权限枚举 —— 所有支持的权限
type Permission string

const (
	PermRead   Permission = "read"
	PermWrite  Permission = "write"
	PermDelete Permission = "delete"
	PermAdmin  Permission = "admin"
	PermExport Permission = "export"
)

// 2. 内存中的权限缓存：groupID -> 拥有的权限集合
var (
	cache     map[int]map[Permission]struct{}
	cacheOnce sync.Once
	initErr   error
)

// Init 主动初始化缓存（如果不调用，第一次 HasPermission 会自动 init）
func Init() error {
	cacheOnce.Do(func() { initErr = loadPermissionGroups() })
	return initErr
}

// 一次性把 PermissionGroup 表读进内存
func loadPermissionGroups() error {
	db := database.DB // 你的 *sql.DB
	if db == nil {
		return fmt.Errorf("database.DB is nil")
	}

	// 这里假设表结构：
	// CREATE TABLE PermissionGroup (
	//   id     INT PRIMARY KEY,
	//   read   BOOLEAN,
	//   write  BOOLEAN,
	//   delete BOOLEAN,
	//   admin  BOOLEAN,
	//   export BOOLEAN
	// );
	rows, err := db.Query(`
		SELECT id, read, write, delete, admin, export
		FROM PermissionGroup`)
	if err != nil {
		return fmt.Errorf("query PermissionGroup: %w", err)
	}
	defer rows.Close()

	tmp := make(map[int]map[Permission]struct{})
	for rows.Next() {
		var (
			id                      int
			read, write, del, adm, exp bool
		)
		if err := rows.Scan(&id, &read, &write, &del, &adm, &exp); err != nil {
			return fmt.Errorf("scan row: %w", err)
		}
		m := make(map[Permission]struct{})
		if read {
			m[PermRead] = struct{}{}
		}
		if write {
			m[PermWrite] = struct{}{}
		}
		if del {
			m[PermDelete] = struct{}{}
		}
		if adm {
			m[PermAdmin] = struct{}{}
		}
		if exp {
			m[PermExport] = struct{}{}
		}
		tmp[id] = m
	}
	if err := rows.Err(); err != nil {
		return err
	}
	cache = tmp
	return nil
}

// HasPermission 判断指定权限组是否拥有「全部」给定权限
// 如果缓存未初始化，会自动 init（懒加载）。
func HasPermission(permGroupId int, perms ...Permission) bool {
	if err := Init(); err != nil {
		// 一旦加载失败，默认拒绝
		return false
	}
	set, ok := cache[permGroupId]
	if !ok {
		// 不存在的权限组
		return false
	}
	for _, p := range perms {
		if _, exist := set[p]; !exist {
			return false
		}
	}
	return true
}