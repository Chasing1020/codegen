/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-11:23
File: config.go
*/

package conf

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"time"
)

var Conf Configuration

type Configuration struct {
	MySQL MySQLConf `json:"mysql" yaml:"mysql"`
	Redis RedisConf `json:"redis" yaml:"redis"`
}

type MySQLConf struct {
	DSN             string        `json:"dsn,omitempty" yaml:"dsn"`
	MaxIdleConn     int           `json:"max_idle_conn,omitempty" yaml:"max_idle_conn"`
	MaxOpenConn     int           `json:"max_open_conn,omitempty" yaml:"max_open_conn"`
	ConnMaxIdleTime time.Duration `json:"conn_max_idle_time,omitempty" yaml:"conn_max_idle_time"`
	ConnMaxLifetime time.Duration `json:"conn_max_lifetime,omitempty" yaml:"conn_max_life_time"`
}

type RedisConf struct {
	Addr         string        `json:"addr,omitempty" yaml:"addr"`
	Password     string        `json:"password,omitempty" yaml:"password"`
	DB           int           `json:"db,omitempty" yaml:"db"`
	MaxRetries   int           `json:"max_retries,omitempty" yaml:"max_retries"`
	ReadTimeout  time.Duration `json:"read_timeout" yaml:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout,omitempty" yaml:"write_timeout"`
}

func init() {
	data, err := os.ReadFile("./conf/config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, &Conf)
	if err != nil {
		panic(err)
	}
	log.Printf("MySQL: %+v",Conf.MySQL)
	log.Printf("Redis: %+v",Conf.Redis)
}
