package config

import (
	"fmt"
	"github.com/gokit/going/db"
	"github.com/gokit/going/logger"
	"github.com/gokit/going/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"log"
)

// Option for configurations
type Option struct {
	Name string `json:"name"` // 应用程序名称
	HTTP struct {
		Host string `json:"host"` // 服务地址
		Port string `json:"port"` // 服务端口
	} `json:"http"`
	Database *db.Config     `json:"mysql"`
	Logger   *logger.Config `json:"logger"`

	Environment string // prod, dev, test
}

// AppConfig is the configs for the whole application
var Instance *Option

// Init is using to initialize the configs
func Init(path string, env string) error {

	var config = viper.New()

	if utils.IsFile(path) {
		config.SetConfigFile(path)
	} else {

		config.SetConfigType("toml")

		if env == "" {
			config.SetConfigName("config")
		} else {
			config.SetConfigName(fmt.Sprintf("config-%s", env))
		}

		config.AddConfigPath(path)
		config.AddConfigPath(".")
		config.AddConfigPath("./config")
		config.AddConfigPath("./configs")
	}

	if err := config.ReadInConfig(); err != nil {
		return err
	}

	log.Printf("loaded config file %s", config.ConfigFileUsed())

	decoderConfig := func(config *mapstructure.DecoderConfig) {
		config.TagName = "json"
	}

	if err := config.Unmarshal(&Instance, decoderConfig); err != nil {
		return err
	}

	Instance.Environment = env

	return nil
}
