package database

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

var appConfigInfo *ConfigModel

type ConfigModel struct {
	Server *ServerModel `yaml:"server"`
}

type ServerModel struct {
	Mode               string `yaml:"mode"`                 // gin运行模式
	Host               string `yaml:"host"`                 // 运行访问域
	Port               string `yaml:"port"`                 // 端口
	EnableConsoleRoute bool   `yaml:"enable_console_route"` // 控制台是否输出路由
}

func ReadAppConfig() (*ConfigModel, error) {
	fPath, _ := os.Getwd()
	fPath = path.Join(fPath, "config/app.yml")
	configData, _err := ioutil.ReadFile(fPath)
	if _err != nil {
		fmt.Printf(" config file read failed: %s", _err)
		os.Exit(-1)
	}
	_err = yaml.Unmarshal(configData, &appConfigInfo)
	if _err != nil {
		fmt.Printf(" config parse failed: %s", _err)
		os.Exit(-1)
	}
	return appConfigInfo, _err
}
