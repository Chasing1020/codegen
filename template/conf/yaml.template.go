/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-16:38
File: conf_yaml_tpl.go
*/

package conf

import "time"

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


var YamlTemplate = `
mysql:
  dsn: {{.MySQL.DSN}}
  max_idle_conn: {{.MySQL.MaxIdleConn}}
  max_open_conn: {{.MySQL.MaxOpenConn}}
  conn_max_idle_time: {{.MySQL.ConnMaxIdleTime}}
  conn_max_life_time: {{.MySQL.ConnMaxLifetime}}

redis:
  addr: {{.Redis.Addr}}
  password: {{.Redis.Password}}
  db: {{.Redis.DB}}
  max_retries: {{.Redis.MaxRetries}}
  read_timeout: {{.Redis.ReadTimeout}}
  write_timeout: {{.Redis.WriteTimeout}}
`
