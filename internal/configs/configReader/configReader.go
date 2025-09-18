package configreader

import (
	"log"
	"sync/atomic"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//#####PUBLIC#####

// 定义配置内容结构体，使包外代码不再依赖配置文件的编写
// 注意：此处结构体开放给外部写死不可改变，请通过改变别名tag来对应实际配置的键名
type DatabaseCfg struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

type InternalAppCfg struct {
	Database DatabaseCfg
}

// Get 并发安全返回最新配置，这是configReader的唯一对外接口
func GetConfig() *InternalAppCfg {
	return atomicCfg.Load().(*InternalAppCfg)
}

//#####PRIVATE#####

// 存放配置的原子容器，局部变量
var atomicCfg atomic.Value

// 转换器结构体，swaper疑似是我自己造的单词
// 方便起见，swaper不要嵌套，直接用下划线表示层级
type swaper struct {
	DatabaseCfg_Host     string `mapstructure:"database.host"`
	DatabaseCfg_Port     int    `mapstructure:"database.port"`
	DatabaseCfg_Username string `mapstructure:"database.username"`
	DatabaseCfg_Password string `mapstructure:"database.password"`
	DatabaseCfg_DBName   string `mapstructure:"database.dbname"`
}

// 局部配置读写器，使用 viper
var configReader = viper.New()

// 初始化配置读写器
func init() {
	log.Print("[INFO] 初始化配置读写器")
	configReader.AddConfigPath(".internal/configs/") //搜索目录
	configReader.SetConfigName("configs")            //配置文件名称
	configReader.SetConfigType("yaml")
	//首次读配置文件
	rcfg_err := configReader.ReadInConfig()
	if rcfg_err != nil {
		log.Fatalf("[FATAL] 无法读取配置文件 错误：%v", rcfg_err)
	}
	if err := swapConfig(configReader); err != nil {
		log.Fatalf("[FATAL][configReader] 首次解析配置失败 错误: %v", err)
	}
	//实现配置文件热加载
	configReader.WatchConfig()
	configReader.OnConfigChange(hotLoadCfg)
}

func hotLoadCfg(e fsnotify.Event) {
	log.Printf("[WARN][configReader] 配置文件变动，开始热加载: %s\n", e.Name)
	if err := swapConfig(configReader); err != nil {
		log.Printf("[ERROR][configReader] 热加载失败，配置未更新: %v\n", err)
	}
}

func swapConfig(v *viper.Viper) error {
	var l swaper
	if err := v.Unmarshal(&l); err != nil {
		return err
	}
	// 转换成对外嵌套模型
	newCfg := &InternalAppCfg{
		Database: DatabaseCfg{
			Host:     l.DatabaseCfg_Host,
			Port:     l.DatabaseCfg_Port,
			Username: l.DatabaseCfg_Username,
			Password: l.DatabaseCfg_Password,
			DBName:   l.DatabaseCfg_DBName,
		},
	}
	atomicCfg.Store(newCfg)
	log.Printf("[INFO][configReader] 配置已更新: %+v\n", *newCfg)
	return nil
}
