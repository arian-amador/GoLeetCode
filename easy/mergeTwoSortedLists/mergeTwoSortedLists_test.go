package mergeTwoSortedLists

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Both Empty", args{nil, nil}, []string{}},
		{"L1 Empty", args{nil, newList([]int{1, 2, 3})}, []string{"1", "2", "3"}},
		{"L2 Empty", args{newList([]int{1, 2, 3}), nil}, []string{"1", "2", "3"}},
		{"Merge", args{newList([]int{1, 2, 3}), newList([]int{2, 3, 4})}, []string{"1", "2", "2", "3", "3", "4"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got.output(), tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got.output(), tt.want)
			}
		})
	}
}
