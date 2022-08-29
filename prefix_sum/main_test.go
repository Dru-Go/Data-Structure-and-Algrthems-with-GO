package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_prefixSum(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test with empty array",
			args: args{
				arr: []int32{},
			},
			want: 0,
		},
		{
			name: "test with valid array",
			args: args{
				arr: []int32{1, 2, 3, 4, 5},
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := prefixSum(tt.args.arr)
			assert.Equal(t, tt.want, got)
		})
	}
}
