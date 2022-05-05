/*
Copyright © 2022 zjc <chasing1020@gmail.com>
Time: 2022/5/5-13:36
File: case_test.go
*/

package config

import (
	"fmt"
	"testing"
)

func TestCase(t *testing.T) {
	fmt.Println(ToSnakeCase("CourseOutline"))
	fmt.Println(ToUpperCamelCase("course_outline"))
}
