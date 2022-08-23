package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_insertNodeAtPosition(t *testing.T) {
	type args struct {
		llist    *SinglyLinkedListNode
		data     int32
		position int32
	}
	tests := []struct {
		name string
		args args
		want *SinglyLinkedListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := insertNodeAtPosition(tt.args.llist, tt.args.data, tt.args.position)
			assert.Equal(t, tt.want, got)
		})
	}
}
