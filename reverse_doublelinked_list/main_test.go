package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverse(t *testing.T) {
	type args struct {
		llist *DoublyLinkedListNode
	}
	var list_1, list_2, list_3 *DoublyLinkedListNode

	list_1 = &DoublyLinkedListNode{
		data: 4,
		prev: nil,
		next: list_2,
	}
	list_2 = &DoublyLinkedListNode{
		data: 5,
		prev: list_1,
		next: list_3,
	}
	list_3 = &DoublyLinkedListNode{
		data: 6,
		prev: list_2,
		next: nil,
	}

	reverseList := &DoublyLinkedListNode{
		data: 6,
		prev: nil,
		next: &DoublyLinkedListNode{
			data: 5,
			prev: list_3,
			next: &DoublyLinkedListNode{
				data: 4,
				prev: list_2,
				next: nil,
			},
		},
	}
	tests := []struct {
		name string
		args args
		want *DoublyLinkedListNode
	}{
		{
			name: "test with empty list",
			args: args{
				llist: list_1,
			},
			want: reverseList,
		},
		{
			name: "test with valid list",
			args: args{
				llist: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverse(tt.args.llist)
			assert.Equal(t, tt.want, got)
		})
	}
}
