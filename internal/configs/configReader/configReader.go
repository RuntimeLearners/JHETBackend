package configreader

import (
	"log"
	"sync"
	"sync/atomic"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//#####CONST#####

const appConfigPath = "../configs"
const appConfigName = "appConfigs"

//#####PUBLIC#####

// 定义配置内容结构体，使包外代码不再依赖配置文件的编写
// 注意：此处结构体开放给外部写死不可改变，请通过改变别名tag来对应实际配置的键名
type DatabaseCfg struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

type InternalAppCfg struct {
	Database DatabaseCfg `mapstructure:"database"`
}

// Get 并发安全返回最新配置，这是configReader的唯一对外接口
var initOnce sync.Once

func GetConfig() InternalAppCfg {
	initOnce.Do(initConfigReader) //首次调用初始化（延迟初始化）
	return atomicCfg.Load().(InternalAppCfg)
}

//#####PRIVATE#####

// 存放配置的原子容器，局部变量
var atomicCfg atomic.Value

// 局部配置读写器，使用 viper
var configReader = viper.New()

// 初始化配置读写器

func initConfigReader() {
	log.Print("[INFO] 载入程序配置")
	configReader.AddConfigPath(appConfigPath) //搜索目录
	configReader.SetConfigName(appConfigName) //配置文件名称
	configReader.SetConfigType("yaml")
	//首次读配置文件
	rcfg_err := configReader.ReadInConfig()

	if rcfg_err != nil {
		log.Fatalf("[FATAL][configReader] 无法读取配置文件 错误：%v", rcfg_err)
	}
	if err := updateConfig(configReader); err != nil {
		log.Fatalf("[FATAL][configReader] 首次解析配置失败 错误: %v", err)
	}
	//实现配置文件热加载
	configReader.WatchConfig()
	configReader.OnConfigChange(hotLoadCfg)
}

func hotLoadCfg(e fsnotify.Event) {
	log.Printf("[WARN][configReader] 配置文件变动，开始热加载: %s\n", e.Name)
	if err := updateConfig(configReader); err != nil {
		log.Printf("[ERROR][configReader] 热加载失败，配置未更新: %v\n", err)
	}
}

func updateConfig(viper *viper.Viper) error {
	var icfg InternalAppCfg
	if err := viper.Unmarshal(&icfg); err != nil {
		return err
	}
	atomicCfg.Store(icfg)
	log.Printf("[INFO][configReader] 程序配置已更新: %+v\n", &icfg)
	return nil
}
