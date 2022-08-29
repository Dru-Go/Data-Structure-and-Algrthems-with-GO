package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_arrayManipulation(t *testing.T) {
	type args struct {
		n       int32
		queries [][]int32
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test with empty array",
			args: args{
				n:       0,
				queries: [][]int32{},
			},
			want: 0,
		},
		{
			name: "test with valid array",
			args: args{
				n:       6,
				queries: [][]int32{{1, 2, 100}, {2, 5, 100}, {3, 4, 100}},
			},
			want: 200,
		},
		{
			name: "test with valid array",
			args: args{
				n:       10,
				queries: [][]int32{{2, 6, 8}, {3, 5, 7}, {1, 8, 1}, {5, 9, 15}},
			},
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := arrayManipulation(tt.args.n, tt.args.queries)
			assert.Equal(t, tt.want, got)
		})
	}
}
