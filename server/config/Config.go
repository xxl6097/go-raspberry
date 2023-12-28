package config

import (
	"github.com/xxl6097/go-glog/glog"
	"github.com/xxl6097/go-http/server/util"
)

type ServerConfig struct {
	Port    int    `yaml:"port"`
	Host    string `yaml:"host"`
	ApiPath string `yaml:"apipath"`
}
type LogViewConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
type HostConfig struct {
	Host string `yaml:"host"`
}

type WebHookConfig struct {
	WX HostConfig `yaml:"wx"`
	DD HostConfig `yaml:"dd"`
}

type Config struct {
	Server  ServerConfig  `yaml:"server"`
	Logview LogViewConfig `yaml:"logview"`
	Webhook WebHookConfig `yaml:"webhook"`
	Token   string        `yaml:"token"`
}

func init() {
	glog.Info("init...")
	util.ParseYaml(conf)
	glog.Infof("%+v", *conf)
}

var conf = &Config{}

func Get() *Config {
	return conf
}
