package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solve(t *testing.T) {
	type args struct {
		arr     []int32
		queries []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "test with empty array",
			args: args{
				arr:     []int32{},
				queries: []int32{},
			},
			want: []int32{},
		},
		{
			name: "test with valid array",
			args: args{
				arr:     []int32{1, 2, 3, 4, 5},
				queries: []int32{1, 2, 3, 4, 5},
			},
			want: []int32{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solve(tt.args.arr, tt.args.queries)
			assert.Equal(t, tt.want, got)
		})
	}
}
