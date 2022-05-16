/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-19:19
File: start.go
*/

package cmd

import (
	"codegen/config"
	"codegen/gen"
	"fmt"
	"os"
	"sync"
)

func Start() {
	createDir()
	genCode()
	//modTidy()
}

var PathList = []string{"/", "/auth", "/conf", "/dal", "/handler", "/model", "/router", "/dal/test"}

func createDir() {
	for _, path := range PathList {
		err := createDirIfNotExists(config.ProjectPath() + "/dist"+path)
		if err != nil {
			panic(err)
		}
	}
}

func createDirIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, 0755); err != nil {
			return err
		}
	}
	return nil
}

func genCode() {
	var wg sync.WaitGroup
	wg.Add(len(gen.Functions))
	for _, f := range gen.Functions {
		go func(f func()) {
			defer func() {
				wg.Done()
				if info := recover(); info != nil {
					fmt.Println(info)
				}
			}()
			f()
		}(f)
	}
	wg.Wait()
	fmt.Println("===== Codegen Success! =====")
	fmt.Println("1. Use `cd dist` to change to generated code directory")
	fmt.Println("2. Use `make fmt` to prepare the environment")
	fmt.Println("3. Use `make run` to start")
}

