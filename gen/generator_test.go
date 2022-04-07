/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/4/7-20:20
File: generator_test.go
*/

package gen

import (
	"testing"
)

func TestGenConf(t *testing.T) {
	Conf()
	Dal()
	Handler()
	Model()
	Router()
	Main()
}
