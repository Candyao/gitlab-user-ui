package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Public struct {
	Gitlab `yaml:"gitlab"`
}

type Gitlab struct {
	Url string         `yaml:"url"`
	Token string       `yaml:"token"`
	EmailDomain string `yaml:"emaildomain"`
}

type configManager struct {
	Public
}

var ConfigManager=&configManager{}

func (this *configManager) Init(configFilePtr *string) {
	this.Public=Public{}
	this.loadConfig(configFilePtr)
}

func (this *configManager) loadConfig(configFilePtr *string) {
	public := viper.New()
	if configFilePtr == nil {
		public.SetConfigFile("./config/config.yaml")
	} else {
		public.SetConfigFile(*configFilePtr)
	}

	public.SetConfigType("yaml")
	if err := public.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
	if err := public.Unmarshal(&this.Public); err != nil {
		fmt.Println(err)
	}
}
