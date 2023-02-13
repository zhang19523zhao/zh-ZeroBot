package config

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

// Configuration 项目配置
type Configuration struct {
	// gtp apikey
	ApiKey string `json:"api_key"`
	// 触发命令,例如触发命令是ai --> @zh /ai 今天天气怎么样
	Command string `json:"command"`
	// 关键词匹配, 只要发送的消息里有关键词就出发
	Regexp string `json:"regexp"`
}

var config *Configuration
var once sync.Once

// LoadConfig 加载配置
func LoadConfig() *Configuration {
	once.Do(func() {
		// 从文件中读取
		config = &Configuration{}
		f, err := os.Open("config.json")
		if err != nil {
			log.Fatalf("open config err: %v", err)
			return
		}
		defer f.Close()
		encoder := json.NewDecoder(f)
		err = encoder.Decode(config)
		if err != nil {
			log.Fatalf("decode config err: %v", err)
			return
		}

		// 如果环境变量有配置，读取环境变量
		ApiKey := os.Getenv("ApiKey")
		Command := os.Getenv("Command")
		Regexp := os.Getenv("Regexp")
		if ApiKey != "" {
			config.ApiKey = ApiKey
		}
		if Command == "" {
			config.Command = Command
		}
		if config.Regexp == "" {
			config.Regexp = Regexp
		}
	})
	return config
}
