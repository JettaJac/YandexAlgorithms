package test

import (
	"main/base"
	"testing"
)

func TestTask(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  string
	}{
		{
			name:  "test1",
			input: []int{1, 1, 1, 1, 1},
			want:  "YES",
		},
		{
			name:  "test2",
			input: []int{2, 2, 2, 1, 1},
			want:  "NO",
		},
		{
			name:  "test3",
			input: []int{5, 5, 14, 5, 6},
			want:  "YES",
		},
		{
			name:  "test4",
			input: []int{5, 5, 14, 5, 5},
			want:  "YES",
		},
		{
			name:  "test5",
			input: []int{5, 15, 14, 5, 14},
			want:  "YES",
		},
		{
			name:  "test5",
			input: []int{5, 15, 14, 5, 5},
			want:  "NO",
		},
		{
			name:  "test6",
			input: []int{5, 15, 1, 5, 1},
			want:  "YES",
		},
		{
			name:  "test6",
			input: []int{15, 5, 1, 5, 1},
			want:  "YES",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := base.BasicFunc(tt.input[0], tt.input[1], tt.input[2], tt.input[3], tt.input[4])
			if res != tt.want {
				t.Errorf("Error result")
			}
			// fmt.Println(res, tt.want)
		})
	}
}
