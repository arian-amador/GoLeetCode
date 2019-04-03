package main

import (
	"reflect"
	"testing"
)

func TestNewList(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Empty slice", args{[]int{}}, ""},
		{"Single digit slice", args{[]int{1}}, "1"},
		{"Small slice", args{[]int{1, 2, 3, 4}}, "1234"},
		{"Large slice", args{[]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}}, "1000000000000000000000000000001"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewList(tt.args.nums); !reflect.DeepEqual(got.debug(), tt.want) {
				t.Errorf("NewList() = %v, want %v", got.debug(), tt.want)
			}
		})
	}
}

func TestAddTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Empty lists", args{NewList([]int{}), NewList([]int{})}, ""},
		{"Single slice", args{NewList([]int{1}), NewList([]int{})}, "1"},
		{"Small slices", args{NewList([]int{1}), NewList([]int{1})}, "2"},
		{"Small slices 2", args{NewList([]int{2, 4, 3}), NewList([]int{5, 6, 4})}, "708"},
		{"Large slices", args{NewList([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}), NewList([]int{5, 6, 4})}, "6640000000000000000000000000001"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got.debug(), tt.want) {
				t.Errorf("AddTwoLists() = %v, want %v", got.debug(), tt.want)
			}
		})
	}
}
