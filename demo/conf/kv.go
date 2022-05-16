/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/5/7-17:50
File: kv.go
*/

package conf

import (
	"crud/consul"
	"encoding/json"
	"gopkg.in/yaml.v2"
	"os"
)

func GetConsulConfig() *Configuration {
	kv := consul.Client.KV()

	var mysql MySQLConf
	var redis RedisConf
	redisConf, _, err := kv.Get("REDIS_CONF", nil)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(redisConf.Value, &redis)

	mysqlConf, _, err := kv.Get("MYSQL_CONF", nil)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(mysqlConf.Value, &mysql)

	return &Configuration{
		MySQL: mysql,
		Redis: redis,
	}
}

func GetLocalConfig() *Configuration {
	data, err := os.ReadFile(ProjectPath() + "/conf/config.yaml")
	if err != nil {
		panic(err)
	}
	var c Configuration
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		panic(err)
	}
	return &c
}