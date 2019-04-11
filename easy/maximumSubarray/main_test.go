package maxSubarray

import "testing"

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Empty", args{[]int{}}, 0},
		{"1 Positive", args{[]int{1}}, 1},
		{"1 Negative", args{[]int{-1}}, -1},
		{"2 Nums", args{[]int{-1, 0}}, 0},
		{"Nums", args{[]int{-1, 0, 23, 12, -102, 12, 1, -1, -2, 0, 10}}, 35},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
