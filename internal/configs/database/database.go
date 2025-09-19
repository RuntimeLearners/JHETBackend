package database

import (
	configreader "JHETBackend/internal/configs/configReader"
	"context"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DataBase *gorm.DB

func Init() error {

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

	// 使用 GORM 打开数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 关闭外键约束以提升迁移速度
	})

	sqlDB, err := db.DB() // 从 *gorm.DB 取出底层 *sql.DB
	if err != nil {
		log.Fatalf("get sql.DB failed: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := sqlDB.PingContext(ctx); err != nil {
		log.Fatalf("ping mysql failed: %v", err)
	}
	log.Println("mysql connect ok")

	// 如果连接失败，返回错误
	if err != nil {
		return fmt.Errorf("database connect failed: %w", err)
	}
	// // 自动迁移数据库结构
	// if err = autoMigrate(db); err != nil {
	// 	return fmt.Errorf("database migrate failed: %w", err)
	// }

	// // 将数据库实例赋值给全局变量 DB
	// DB = db
	return nil
}
