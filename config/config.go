package config

import (
	"fmt"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/klog"
	"os"
)

type KeyName string

const (
	ServerName KeyName = "server_name"
	ServerHost KeyName = "server_host"
	ServerPort KeyName = "server_port"
)

var keyMap map[KeyName]interface{}

type Config struct {
	Server Server
}

type Server struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func init() {
	var config Config
	/*dir, err := os.Getwd()
	if err != nil {
		klog.Fatal(err)
		return
	}*/
	yamlFile, err := ioutil.ReadFile("./config/application.yaml")
	if err != nil {
		klog.Fatal(err)
		return
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		klog.Fatal(err)
		return
	}
	fmt.Println(os.Getwd())
	keyMap = make(map[KeyName]interface{})
	keyMap[ServerName] = config.Server.Name
	keyMap[ServerHost] = config.Server.Host
	keyMap[ServerPort] = config.Server.Port
}

func GetString(keyName KeyName) string {
	return keyMap[keyName].(string)
}

func GetInt(keyName KeyName) int {
	return keyMap[keyName].(int)
}
