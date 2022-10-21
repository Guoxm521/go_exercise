package database

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

var appConfigInfo *ConfigModel

type ConfigModel struct {
	Server *ServerModel `yaml:"server"`
	MySql  *MysqlModel  `yaml:"mysql"`
}

type ServerModel struct {
	Mode               string `yaml:"mode"`                 // gin运行模式
	Host               string `yaml:"host"`                 // 运行访问域
	Port               string `yaml:"port"`                 // 端口
	EnableConsoleRoute bool   `yaml:"enable_console_route"` // 控制台是否输出路由
	AppID              string `yaml:"app_id"`               // 应用ID
	AppKey             string `yaml:"app_key"`              // 应用key
}

func ReadAppConfig() (*ConfigModel, error) {
	fPath, _ := os.Getwd()
	fPath = path.Join(fPath, "config")
	configPath := flag.String("c", fPath, "config file path")
	flag.Parse()
	err := loadConfigInfo(*configPath, "app.yml", &appConfigInfo)
	return appConfigInfo, err
}

func loadConfigInfo(configPath string, yml string, structObject interface{}) (err error) {
	var (
		filePath string
		wr       string
	)
	if configPath == "" {
		wr, _ = os.Getwd()
		wr = path.Join(wr, "config")
	} else {
		wr = path.Join(wr, configPath)
	}
	filePath = path.Join(wr, yml)
	configData, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf(" config file read failed: %s", err)
		os.Exit(-1)
	}
	err = yaml.Unmarshal(configData, structObject)
	if err != nil {
		fmt.Printf(" config parse failed: %s", err)
		os.Exit(-1)
	}
	return nil
}
