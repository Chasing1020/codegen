/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-16:57
File: init.go
*/

package config

import (
	"time"
)

type Configuration struct {
	*Module `yaml:"module"`
	*Schema `yaml:"schema"`
}

type Module struct {
	Package string    `yaml:"package"`
	Author  string    `yaml:"author"`
	Email   string    `yaml:"email"`
	Time    time.Time `yaml:"time"`
	MySQL   MySQLConf `yaml:"mysql"`
	Redis   RedisConf `yaml:"redis"`
}

type Schema struct {
	Tables []*Table `yaml:"tables"`
}

type Table struct {
	Name    string    `yaml:"name"`
	Tag     string    `yaml:"Tag"`
	Columns []*Column `yaml:"columns"`
}

type Column struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
	Tag  string `yaml:"tag"`
}

type MySQLConf struct {
	DSN             string `json:"dsn,omitempty" yaml:"dsn"`
	MaxIdleConn     int    `json:"max_idle_conn,omitempty" yaml:"max_idle_conn"`
	MaxOpenConn     int    `json:"max_open_conn,omitempty" yaml:"max_open_conn"`
	ConnMaxIdleTime string `json:"conn_max_idle_time,omitempty" yaml:"conn_max_idle_time"`
	ConnMaxLifetime string `json:"conn_max_lifetime,omitempty" yaml:"conn_max_life_time"`
}

type RedisConf struct {
	Addr         string `json:"addr,omitempty" yaml:"addr"`
	Password     string `json:"password,omitempty" yaml:"password"`
	DB           int    `json:"db,omitempty" yaml:"db"`
	MaxRetries   int    `json:"max_retries,omitempty" yaml:"max_retries"`
	ReadTimeout  string `json:"read_timeout" yaml:"read_timeout"`
	WriteTimeout string `json:"write_timeout,omitempty" yaml:"write_timeout"`
}
