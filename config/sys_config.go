/*
 * @Description: System parameters initialize the configuration file
 * @Author: Young (hao_youngg@163.com)
 * @LastEditors: Young (hao_youngg@163.com)
 * @Date: 2019-04-30 11:35:20
 * @LastEditTime: 2019-04-30 16:43:06
 */
package config

import (
	"base_core/log"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	Redis RedisConfig // Redis配置
	Psql  PsqlConfig  // PostgreSQL配置
	Env   EnvConfig   // 系统启动环境配置
	Jwt   JwtConfig   // JWT 密码配置
}

type PsqlConfig struct {
	Host    string `json:"host"`
	Port    string `json:"port"`
	DB      string `json:"db"`
	User    string `json:"user"`
	Pwd     string `json:"pwd"`
	Sslmode string `json:"sslmode"`
	IsShow  bool   `json:"isShow"`
}
type RedisConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Db   int32  `json:"db"`
	Pwd  string `json:"pwd"`
}

type EnvConfig struct {
	Env  string `json:"env"`
	Port string `json:"port"`
	Host string `json:"host"`
}

type JwtConfig struct {
	Pwd []byte `json:"pwd"`
}

func ConfInit(path string) {
	Conf = NewConfig(path)
}

func NewConfig(config string) *Config {
	absPath, _ := filepath.Abs(config)
	config_str, err := ioutil.ReadFile(absPath)
	if err != nil {
		log.DaoLogger.Error("NewConfig: ", err)
	}

	config_obj := Config{}
	e := json.Unmarshal(config_str, &config_obj)
	if e != nil {
		fmt.Println(e.Error())
		fmt.Println("配置文件 json 解析失败")
		os.Exit(0)
	}

	return &config_obj
}
