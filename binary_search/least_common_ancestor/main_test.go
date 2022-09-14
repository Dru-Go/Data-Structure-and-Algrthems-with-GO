package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_leastCommonAncestorInBST(t *testing.T) {
	type args struct {
		root *Node
		n1   int32
		n2   int32
	}
	root := Node{100, nil, nil}
	root.left = &Node{50, nil, nil}
	root.right = &Node{150, nil, nil}
	root.left.left = &Node{20, nil, nil}
	root.left.left.left = &Node{10, nil, nil}
	root.left.left.right = &Node{30, nil, nil}
	root.left.right = &Node{80, nil, nil}
	root.right.left = &Node{110, nil, nil}
	root.right.right = &Node{200, nil, nil}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "test with empty root",
			args: args{
				root: &Node{},
				n1:   10,
				n2:   15,
			},
			want: -1,
		},
		{
			name: "test with a valid root #1",
			args: args{
				root: &root,
				n1:   10,
				n2:   15,
			},
			want: 10,
		},
		{
			name: "test with a valid root #2",
			args: args{
				root: &root,
				n1:   10,
				n2:   80,
			},
			want: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := leastCommonAncestorInBST(tt.args.root, tt.args.n1, tt.args.n2)
			assert.Equal(t, tt.want, got)

		})
	}
}
