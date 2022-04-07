/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-11:13
File: main.go
*/

package main

import (
	"codegen/router"
)

func main() {
	engine := router.InitEngine()
	err := engine.Run(":8080")
	if err != nil {
		panic(err)
	}

}
