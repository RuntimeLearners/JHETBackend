package configreader

import (
	"log"

	"github.com/spf13/viper"
)

// 全局配置读写器，使用 viper
var ConfigReader = viper.New()

// 初始化配置读写器
func init() {
	log.Print("[INFO] 初始化配置读写器")
	ConfigReader.AddConfigPath(".internal/configs/") //搜索目录
	ConfigReader.SetConfigName("configs")            //配置文件名称
	ConfigReader.SetConfigType("yaml")
	//首次读配置文件
	rcfg_err := ConfigReader.ReadInConfig()
	if rcfg_err != nil {
		log.Fatalf("[FATAL] 无法读取配置文件 错误：%v", rcfg_err)
	}

}
