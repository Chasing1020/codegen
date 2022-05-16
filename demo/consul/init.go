/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/5/7-16:45
File: init.go
*/

package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

const (
	LocalIP     = "127.0.0.1"
	ServicePort = 8080
)

var Client *api.Client

func init() {
	config := api.DefaultConfig()
	var err error
	Client, err = api.NewClient(config)
	if err != nil {
		panic(err)
	}

	registerHealthChecker()
}

func registerHealthChecker() {
	registration := &api.AgentServiceRegistration{
		Name:    "crud",
		Port:    ServicePort,
		Tags:    []string{"crud_test_service"},
		Address: LocalIP,
	}

	actuators := []string{"", "/redis", "/mysql"}
	for _, actuator := range actuators {
		registration.Checks = append(registration.Checks, &api.AgentServiceCheck{
			HTTP:                           fmt.Sprintf("http://%s:%d/actuator/health%s", LocalIP, ServicePort, actuator),
			Timeout:                        "10s",
			Interval:                       "10s",
			DeregisterCriticalServiceAfter: "30s",
		})
	}

	err := Client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
}

