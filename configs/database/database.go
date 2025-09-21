package database

import (
	configreader "JHETBackend/configs/configReader"
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DataBase *gorm.DB

var initDBOnce sync.Once

func init() { // 立即加载数据库，根据需要可以删除
	initDBOnce.Do(initDatabase)
}

func initDatabase() {

	//TODO: 判断连接成功和报错的逻辑还要再改，研究下自动迁移相关的设置

	// 抄了一点4UOnline-Go的代码 折乙你不会生气吧（

	// 从配置中获取数据库连接所需的参数
	user := configreader.GetConfig().Database.Username // 数据库用户名
	pass := configreader.GetConfig().Database.Password // 数据库密码
	host := configreader.GetConfig().Database.Host     // 数据库主机
	port := configreader.GetConfig().Database.Port     // 数据库端口
	name := configreader.GetConfig().Database.DBName   // 数据库名称

	// 构建数据源名称 (DSN)
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name)

	log.Printf("[INFO][DB] 尝试连接到数据库, dsn=%v", dsn)
	var dbtmp *gorm.DB
	var err error
	// 使用 GORM 打开数据库连接
	dbtmp, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 关闭外键约束以提升迁移速度
	})
	if err != nil {
		log.Panicf("[FATAL] Open database failed. Err: %v", err)
	}
	// // 自动迁移数据库结构
	// if err = autoMigrate(db); err != nil {
	// 	return fmt.Errorf("database migrate failed: %w", err)
	// }
	log.Printf("[INFO][DB] 数据库连接成功, 当前数据库: %s", name)
	//连接成功传递到全局指针
	DataBase = dbtmp
}
