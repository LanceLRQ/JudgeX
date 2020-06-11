package core

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type ServerConfiguration struct {
	// 服务器配置
	Server ScServer `yaml:"server" json:"server"`
	// 调试模式
	DebugMode bool `yaml:"debug" json:"debug"`
	// 工作目录
	WorkRoot string `yaml:"work_root" json:"work_root"`
}

type ScServer struct {
	// 监听IP
	Listen string `yaml:"listen" json:"listen"`
	// 监听端口
	Port int `yaml:"port" json:"port,int"`
}

var Settings ServerConfiguration

func LoadServerConfig(configFile string) error {
	if yamlFile, err := ioutil.ReadFile(configFile); err == nil {
		if err = yaml.Unmarshal(yamlFile, &Settings); err != nil {
			return fmt.Errorf("failed to unmarshal configuration: %s", err.Error())
		}
	} else {
		return fmt.Errorf("failed to load configuration: %s", err.Error())
	}
	return nil
}