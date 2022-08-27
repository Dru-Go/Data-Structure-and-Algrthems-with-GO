package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxSubarray(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "test with empty array",
			args: args{
				arr: []int32{},
			},
			want: []int32{},
		},
		{
			name: "test with valid array",
			args: args{
				arr: []int32{2, -1, 2, 3, 4, -5},
			},
			want: []int32{10, 11},
		},
		{
			name: "test with valid array",
			args: args{
				arr: []int32{-2, -3, -1, -4, -6},
			},
			want: []int32{-1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSubarray(tt.args.arr)
			assert.Equal(t, tt.want, got)
		})
	}
}
