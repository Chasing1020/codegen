/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/5/7-17:45
File: kv_test.go
*/

package conf

import (
	"crud/consul"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/consul/api"
	"testing"
)

func TestGetConfig(t *testing.T){
	// Get a handle to the KV API
	kv := consul.Client.KV()

	var redis RedisConf
	redisConf, _, err := kv.Get("REDIS_CONF", nil)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(redisConf.Value, &redis)
	fmt.Printf("redisConf.Value: %+v\n", redis)

	var mysql MySQLConf
	mysqlConf, _, err := kv.Get("MYSQL_CONF", nil)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(mysqlConf.Value, &mysql)
	fmt.Printf("mysqlConf.Value: %+v\n", mysql)
}

func TestSetConfig(t *testing.T) {
	kv := consul.Client.KV()

	redis, _ := json.Marshal(Conf.Redis)
	_, err := kv.Put(&api.KVPair{Key: "REDIS_CONF", Value: redis}, nil)
	if err != nil {
		panic(err)
	}

	mysql, _ := json.Marshal(Conf.MySQL)
	_, err = kv.Put(&api.KVPair{Key: "MYSQL_CONF", Value: mysql}, nil)
	if err != nil {
		panic(err)
	}
}