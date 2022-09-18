package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pairs(t *testing.T) {
	type args struct {
		k   int32
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "test with empty array",
			args: args{
				k:   0,
				arr: []int32{},
			},
			want: 0,
		},
		{
			name: "test with valid array",
			args: args{
				k:   2,
				arr: []int32{1, 5, 3, 4, 2},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pairs(tt.args.k, tt.args.arr)
			assert.Equal(t, tt.want, got)
		})
	}
}
