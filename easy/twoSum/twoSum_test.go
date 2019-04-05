package twoSum

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Sorted", args{[]int{2, 7, 11, 15}, 9}, []int{0, 1}},
		{"Unsorted", args{[]int{15, 7, 2, 11}, 9}, []int{1, 2}},
		{"Multiple Correct", args{[]int{15, 7, 2, 3, 4, 1, 8, 6, 11}, 10}, []int{1, 3}},
		{"Empty result", args{[]int{15, 7, 2, 11}, 100}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bruteTwoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_onePassHash(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Sorted", args{[]int{2, 7, 11, 15}, 9}, []int{0, 1}},
		{"Unsorted", args{[]int{15, 7, 2, 11}, 9}, []int{1, 2}},
		{"Multiple Correct", args{[]int{15, 7, 2, 3, 4, 1, 8, 6, 11}, 10}, []int{1, 3}},
		{"Empty result", args{[]int{15, 7, 2, 11}, 100}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := onePassHash(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("onePassHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortedTwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Sorted", args{[]int{2, 7, 11, 15}, 9}, []int{0, 1}},
		{"Unsorted", args{[]int{15, 7, 2, 11}, 9}, []int{0, 1}},
		{"Multiple Correct", args{[]int{15, 7, 2, 3, 4, 1, 8, 6, 11}, 10}, []int{1, 6}},
		{"Empty result", args{[]int{15, 7, 2, 11}, 100}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedTwoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedTwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
