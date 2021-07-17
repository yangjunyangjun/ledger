package conf

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

var Config YamlConfig

type YamlConfig struct {
	DataBase DataBase `yaml:"database"`
	Redis    Redis    `yaml:"redis"`
	RabbitMq RabbitMq `yaml:"rabbitmq"`
	Log      Log      `yaml:"log"`
	Jwt      Jwt      `yaml:"jwt"`
	Email    Email    `yaml:"email"`
	IconPath string   `yaml:"iconfilepath"`
}

type DataBase struct {
	Ip       string `yaml:"ip"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Db       string `yaml:"db"`
	TimeOut  int64  `yaml:"timeout"`
}

type Redis struct {
	Ip       string `yaml:"ip"`
	Password string `yaml:"password"`
}

type RabbitMq struct {
	Ip       string `yaml:"ip"`
	Port     int64  `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Log struct {
	FIleName string `yaml:"filename"`
	MaxSize  int64  `yaml:"maxsize"`
}

type Jwt struct {
	Secret string `yaml:"secret"`
}

type Email struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func init() {
	input := flag.String("env", "sit", "")
	//flag.Parse()

	pwd, _ := os.Getwd() //获取当前运行目录

	//以项目名分割,取出项目根路径
	pathList := strings.SplitAfter(pwd, "ledger")
	path := pathList[0]

	//配置文件路径
	settingPath := fmt.Sprintf("%s/settings/%s.yaml", path, *input)
	c, err := ioutil.ReadFile(settingPath)
	if err != nil {
		fmt.Println("read setting file err")
		os.Exit(1)
	}
	if err := yaml.Unmarshal(c, &Config); err != nil {
		fmt.Println("yaml unmarshal err")
		os.Exit(1)
	}
}
