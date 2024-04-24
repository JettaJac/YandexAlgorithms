package test

import (
	"main/main2"
	"reflect"
	"testing"
)

func TestTask(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  string
		// testp bool
	}{
		{
			name:  "test1",
			input: []int{1, 1, 2, 1, 1},
			want:  "1 1",
		},
		{
			name:  "test2",
			input: []int{2, 1, 1, 1, 1},
			want:  "0 1",
		},
		{
			name:  "test3",
			input: []int{89, 20, 44, 1, 11},
			want:  "2 3",
		},
		{
			name:  "test4",
			input: []int{89, 20, 40, 1, 11},
			want:  "-1 -1",
		},
		{
			name:  "test5",
			input: []int{1, 1, 10, 2, 1},
			want:  "1 1",
		},
		{
			name:  "test6",
			input: []int{7, 1, 10, 2, 1},
			want:  "0 1",
		},
		{
			name:  "test7",
			input: []int{8, 4, 10, 2, 3},
			want:  "-1 -1",
		},
		{
			name:  "test8",
			input: []int{12, 6, 14, 2, 3},
			want:  "-1 -1",
		},
		{
			name:  "test9",
			input: []int{7, 4, 11, 1, 2},
			want:  "1 2",
		},
		{
			name:  "test10",
			input: []int{15, 5, 20, 3, 4},
			want:  "-1 -1",
		},
		{
			name:  "test11",
			input: []int{12, 3, 18, 2, 3},
			want:  "2 1",
		},
		{
			name:  "test12",
			input: []int{9, 3, 18, 2, 3},
			want:  "1 3",
		},
		{
			name:  "test13",
			input: []int{10, 3, 18, 2, 3},
			want:  "2 1",
		},
		{
			name:  "test14",
			input: []int{100, 1, 2, 1, 1},
			want:  "0 1",
		},
		{
			name:  "test15",
			input: []int{100, 2, 2, 1, 1},
			want:  "0 0",
		},
		{
			name:  "test16",
			input: []int{2, 1, 2, 1, 1},
			want:  "1 1",
		},
		{
			name:  "test17",
			input: []int{45, 2, 45, 2, 3},
			want:  "-1 -1",
		},
		{
			name:  "test18",
			input: []int{43, 20, 43, 1, 11},
			want:  "1 11",
		},
		{
			name:  "test19",
			input: []int{18, 3, 18, 2, 3},
			want:  "2 3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := main2.BaseFunc(tt.input[0], tt.input[1], tt.input[2], tt.input[3], tt.input[4])
			// got := coins.MinCoins(tt.val, tt.coins)
			// got2 := coins.MinCoins2(tt.val, tt.coins)
			if res != tt.want {
				t.Errorf("Error result")

			}

			// fmt.Println(res, "  ", tt.want)
		})

		// if res != tt.expected {
		// 	t.Errorf("Got \"%v\", expected \"%v\"", res, tt.expected)
		// }
	}
}
