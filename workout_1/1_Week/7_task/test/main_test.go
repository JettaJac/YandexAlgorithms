package test

import (
	"fmt"
	"main/base"
	"testing"
)

func TestTask(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "test1",
			input: []int{10, 5, 2},
			want:  4,
		},
		{
			name:  "test2",
			input: []int{13, 5, 3},
			want:  3,
		},
		{
			name:  "test3",
			input: []int{14, 5, 3},
			want:  4,
		},
		{
			name:  "test4",
			input: []int{200, 100, 10},
			want:  20,
		},
		{
			name:  "test5",
			input: []int{200, 4, 3},
			want:  66,
		},
		{
			name:  "test6",
			input: []int{1, 4, 3},
			want:  0,
		},
		{
			name:  "test7",
			input: []int{99, 4, 3},
			want:  32,
		},
		{
			name:  "test8",
			input: []int{200, 1, 1},
			want:  200,
		},
		{
			name:  "test9",
			input: []int{200, 3, 2},
			want:  99,
		},
		{
			name:  "test9",
			input: []int{2000, 3, 2},
			want:  999,
		},
		{
			name:  "test10",
			input: []int{2000, 0, 2},
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := base.BasicFunc(tt.input[0], tt.input[1], tt.input[2])
			if res != tt.want {
				t.Errorf("Error result")
			}
			fmt.Println(res, tt.want)
		})
	}
}
