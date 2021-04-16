package q2

import (
	"reflect"
	"testing"
)

func Test_add(t *testing.T) {
	got := addTwoNumbers(&ListNode{2, &ListNode{4, &ListNode{3, &ListNode{1, nil}}}}, &ListNode{5, &ListNode{6, &ListNode{8, nil}}})
	for got != nil {
		t.Logf("%v\n", got)
		got = got.Next
	}
}

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"test sigle", args{&ListNode{7, nil}, &ListNode{7, nil}}, &ListNode{4, &ListNode{1, nil}}},
		{
			"test multi",
			args{
				&ListNode{2, &ListNode{4, &ListNode{3, &ListNode{1, nil}}}},
				&ListNode{5, &ListNode{6, &ListNode{8, nil}}},
			},
			&ListNode{7, &ListNode{0, &ListNode{2, &ListNode{2, nil}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
