/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-16:32
File: conf_tpl.go
*/

package conf

var Template = `// Copyright © 2022 {{.Author}} <{{.Email}}>
// Time: {{ .Time.Format "2006-01-02T15:04:05Z07:00" }}
// File: conf.go

// Package conf will be used to unmarshal yaml
package conf

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"path"
	"time"
)

var Conf Configuration

type Configuration struct {
	MySQL MySQLConf `+"`"+`json:"mysql" yaml:"mysql"`+"`"+`
	Redis RedisConf `+"`"+`json:"redis" yaml:"redis"`+"`"+`
}

type MySQLConf struct {
	DSN             string        `+"`"+`json:"dsn,omitempty" yaml:"dsn"`+"`"+`
	MaxIdleConn     int           `+"`"+`json:"max_idle_conn,omitempty" yaml:"max_idle_conn"`+"`"+`
	MaxOpenConn     int           `+"`"+`json:"max_open_conn,omitempty" yaml:"max_open_conn"`+"`"+`
	ConnMaxIdleTime time.Duration `+"`"+`json:"conn_max_idle_time,omitempty" yaml:"conn_max_idle_time"`+"`"+`
	ConnMaxLifetime time.Duration `+"`"+`json:"conn_max_lifetime,omitempty" yaml:"conn_max_life_time"`+"`"+`
}

type RedisConf struct {
	Addr         string        `+"`"+`json:"addr,omitempty" yaml:"addr"`+"`"+`
	Password     string        `+"`"+`json:"password,omitempty" yaml:"password"`+"`"+`
	DB           int           `+"`"+`json:"db,omitempty" yaml:"db"`+"`"+`
	MaxRetries   int           `+"`"+`json:"max_retries,omitempty" yaml:"max_retries"`+"`"+`
	ReadTimeout  time.Duration `+"`"+`json:"read_timeout" yaml:"read_timeout"`+"`"+`
	WriteTimeout time.Duration `+"`"+`json:"write_timeout,omitempty" yaml:"write_timeout"`+"`"+`
}

func init() {
	data, err := os.ReadFile(ProjectPath() + "/conf/config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, &Conf)
	if err != nil {
		panic(err)
	}
	log.Printf("MySQL: %+v", Conf.MySQL)
	log.Printf("Redis: %+v", Conf.Redis)
}

var wd string

// ProjectPath returns the path to the project
func ProjectPath() string {
	if wd != "" {
		return wd
	}
	var err error
	wd, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	for i := 0; i < 5; i++ {
		if _, err := os.Stat(wd + "/conf"); os.IsNotExist(err) {
			wd = path.Join(wd, "/..")
		} else {
			break
		}
	}
	return wd
}`