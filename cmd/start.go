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
	"os/exec"
	"sync"
)

func Start() {
	createDir()
	genCode()
	modTidy()
}

var PathList = []string{"/dist/", "/dist/auth", "/dist/conf", "/dist/dal", "/dist/handler", "/dist/model", "/dist/router"}

func createDir() {
	for _, path := range PathList {
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
	fmt.Println("code generation finished")
}

func modTidy() {
	err := os.Chdir(config.ProjectPath() + "/dist")
	if err != nil {
		panic(err)
	}
	err = exec.Command("go", "mod", "tidy").Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("go mod tidy finished")

	for _, path := range []string{"conf", "auth", "dal", "handler", "model", "router"} {
		err = exec.Command("go", "fmt", "./"+path).Run()
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("go source code format finished")

	fmt.Println("===== Codegen Success! =====")

	fmt.Println("use command `cd dist && go run main.go` to start")
	fmt.Println("use `go get -u github.com/swaggo/swag/cmd/swag`")
	fmt.Println("(1.16 or newer)`go install github.com/swaggo/swag/cmd/swag@latest`")
	fmt.Println("use command `swag fmt && swag init` to generate documentation")
}
