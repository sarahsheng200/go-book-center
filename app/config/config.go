package config

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Yaml struct {
	Server `yaml:"server"`
	Redis  `yaml:"redis"`
	Mysql  `yaml:"mysql"`
	//Mongo   `yaml:"mongo"`
	Session `yaml:"session"`
	Log     `yaml:"log"`
	//Url     `yaml:"url"`
	//File    `yaml:"file"`
	//Oss     `yaml:"oss"`
}

type Server struct {
	Port      int    `yaml:"port"`
	Mode      string `yaml:"mode"`
	UseRedis  bool   `yaml:"useRedis"`
	UseMysql  bool   `yaml:"useMysql"`
	UrlPrefix string `yaml:"urlPrefix"`
}

type Mysql struct {
	Path          string `yaml:"path"`
	Database      string `yaml:"database"`
	Username      string `yaml:"username"`
	Password      string `yaml:"password"`
	Config        string `yaml:"config"`
	IsLogMode     bool   `yaml:"isLogMode"`
	SingularTable bool   `yaml:"singularTable"`
}

type Log struct {
	IsDebug  bool   `yaml:"isDebug"`
	FileName string `yaml:"fileName"`
	DirName  string `yaml:"dirName"`
	MaxAge   int    `yaml:"maxAge"`
}

type Redis struct {
	Addr     string `yaml:"Addr"`
	Password string `yaml:"password"`
	Db       int    `yaml:"db"`
}

type Session struct {
	StoreKey   string `yaml:"storeKey"`
	Name       string `yaml:"name"`
	SessionKey string `yaml:"sessionKey"`
	Size       int    `yaml:"size"`
	MaxAge     int    `yaml:"maxAge"`
	Path       string `yaml:"path"`
	Domain     string `yaml:"domain"`
	HttpOnly   bool   `yaml:"httpOnly"`
}

var Conf *Yaml

const defaultConfig = "config.yaml"

func init() {
	c := &Yaml{}
	configFile := flag.String("c", defaultConfig, "config file")
	flag.Parse()
	yamlFile, err := ioutil.ReadFile(*configFile)
	if err != nil {
		panic(fmt.Errorf("get yamlFile error: %s", err))
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal yamlFile error: %s", err)
	}
	Conf = c
}
