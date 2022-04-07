/*
Copyright © 2022 Chasing1020 <chasing1020@gmail.com>
Time: 2022-04-07 23:11:10.330178 +0800 CST m=+0.001950585
File: response.go
*/

package main

import (
	"crud/router"
)

// @title crud
// @version 1.0
// @description crud Project
// @contact.name Chasing1020
// @contact.url https://github.com/Chasing1020
// @contact.email chasing1020@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	engine := router.InitEngine()
	err := engine.Run(":8080")
	if err != nil {
		panic(err)
	}
}