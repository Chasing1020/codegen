/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-19:19
File: start.go
*/

package cmd

import (
	"codegen/config"
	"codegen/gen"
	"log"
	"os"
	"os/exec"
	"sync"
)

func Start() {
	createDir()
	genCode()
	modTidy()
}

func createDir() {
	pathList := []string{"/dist/", "/dist/conf", "/dist/dal", "/dist/handler", "/dist/model", "/dist/router"}
	for _, path := range pathList {
		err := createDirIfNotExists(config.ProjectPath() + path)
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
	for _, f := range gen.Functions {
		wg.Add(1)
		go func(f func()) {
			defer func() {
				wg.Done()
				if info := recover(); info != nil {
					log.Println(info)
				}
			}()
			f()
		}(f)
	}
	wg.Wait()
	log.Println("code generation finished")
}

func modTidy() {
	err := os.Chdir(config.ProjectPath() + "/dist")
	if err != nil {
		panic(err)
	}
	cmd := exec.Command("go", "mod", "tidy")
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
	log.Println("go mod tidy finished")
	log.Println("use command `cd dist && go run main.go` to start")
}
