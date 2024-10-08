package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
)

var config *ChatBotConfig

// ChatBotConfig 配置
type ChatBotConfig struct {
	Mod      string `yaml:"mod"`
	LogLevel int    `yaml:"logLevel"`
	Master   string `yaml:"master"`
}

func LoadConfig() *ChatBotConfig {
	vconfig := viper.New()

	//表示 先预加载匹配的环境变量
	vconfig.AutomaticEnv()
	//设置环境变量分割符，将点号和横杠替换为下划线
	vconfig.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	// 设置读取的配置文件
	vconfig.SetConfigName("config")
	// 添加读取的配置文件路径
	vconfig.AddConfigPath(".")
	// 读取文件类型
	vconfig.SetConfigType("yaml")

	err := vconfig.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	if err := vconfig.Unmarshal(&config); err != nil {
		log.Panicln("\"unmarshal cng file fail " + err.Error())
	}
	// 获取所有环境变量
	return config
}
