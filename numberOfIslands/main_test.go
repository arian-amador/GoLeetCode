package main

import "testing"

func Test_numIslands(t *testing.T) {
	type args struct {
		grid [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Empty", args{[][]byte{}}, 0},
		{"0 Islands", args{[][]byte{{'0', '0', '0', '0', '0'}, {'0', '0', '0', '0', '0'}, {'0', '0', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}}, 0},
		{"1 Islands", args{[][]byte{{'0', '0', '0', '0', '0'}, {'0', '0', '1', '1', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '0', '0'}}}, 1},
		{"2 Islands", args{[][]byte{{'0', '1', '0', '0', '0'}, {'0', '0', '1', '1', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '0', '0'}}}, 2},
		{"3 Islands", args{[][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}}, 3},
		{"I shape", args{[][]byte{{'1', '1', '1'}, {'0', '1', '0'}, {'1', '1', '1'}}}, 1},
		{"Left Bound", args{[][]byte{{'1', '0', '0', '0', '0'}, {'1', '0', '0', '0', '0'}, {'0', '0', '0', '0', '0'}, {'1', '0', '0', '0', '0'}}}, 2},
		{"Full", args{[][]byte{{'1', '1', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '1', '1', '1', '1'}}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands(tt.args.grid); got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}
