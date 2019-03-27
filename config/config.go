package config

import (
	"gin/log"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var configFile []byte

type ChannelConfig struct {
	Channel Channel `yaml:"channel"`
}

type Channel struct {
	EmayReminderConfig EmayReminder `yaml:"emayReminder"`
	GuoduConfig        Guodu        `yaml:"guodu"`
}

type EmayReminder struct {
	UserId    string `yaml:"userId"`
	UserPws   string `yaml:"userPws"`
	Url       string `yaml:"url"`
	Threshold string `yaml:"threshold"`
}

type Guodu struct {
	UserId    string `yaml:"userId"`
	UserPws   string `yaml:"userPws"`
	Url       string `yaml:"url"`
	KeyStr    string `yaml:"keyStr"`
	Threshold string `yaml:"threshold"`
}

func GetConfig() (e *ChannelConfig, err error) {
	err = yaml.Unmarshal(configFile, &e)
	return e, err
}

func init() {
	var err error
	if gin.Mode() == gin.ReleaseMode {
		configFile, err = ioutil.ReadFile("./conf_prod")
		if err != nil {
			log.Warn("read file conf_prod failed", err)
		}
	} else if gin.Mode() == gin.TestMode {
		configFile, err = ioutil.ReadFile("./conf_test")
		if err != nil {
			log.Warn("read file conf_prod failed", err)
		}
	}
	configFile, err = ioutil.ReadFile("conf_dev.yaml")
	if err != nil {
		log.Warn("read file conf_dev failed", err)
	}
}
