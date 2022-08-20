package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverse(t *testing.T) {
	type args struct {
		llist *SinglyLinkedListNode
	}
	list := &SinglyLinkedListNode{
		data: 4,
		next: &SinglyLinkedListNode{
			data: 5,
			next: &SinglyLinkedListNode{
				data: 6,
				next: nil,
			},
		},
	}
	reverseList := &SinglyLinkedListNode{
		data: 6,
		next: &SinglyLinkedListNode{
			data: 5,
			next: &SinglyLinkedListNode{
				data: 4,
				next: nil,
			},
		},
	}
	tests := []struct {
		name string
		args args
		want *SinglyLinkedListNode
	}{
		{
			name: "testing with empty list",
			args: args{
				llist: nil,
			},
			want: nil,
		},
		{
			name: "testing with valid list",
			args: args{
				llist: list,
			},
			want: reverseList,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverse(tt.args.llist)
			assert.Equal(t, tt.want, got)
		})
	}
}
