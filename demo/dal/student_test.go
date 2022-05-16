// Copyright © 2022 Chasing1020 <chasing1020@gmail.com>
// Time: 2022/5/16T11:31:15+08:00
// File: student_test.go

package dal

import (
	"context"
	"crud/model"
	"testing"
)

func TestMockStudent(t *testing.T) {
	ctx := context.Background()

	students := make([]*model.Student, 20)
	for _, student := range students {
		student = &model.Student{
			// TODO: fill all fields
		}
		CreateStudent(ctx, student)
	}
}

