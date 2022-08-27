package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_equalStacks(t *testing.T) {
	type args struct {
		h1 []int32
		h2 []int32
		h3 []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "Test with empty arrays",
			args: args{
				h1: []int32{},
				h2: []int32{},
				h3: []int32{},
			},
			want: 0,
		},
		{
			name: "Test with arrays",
			args: args{
				h1: []int32{3, 2, 1, 1, 1},
				h2: []int32{4, 3, 2},
				h3: []int32{1, 1, 4, 1},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := equalStacks(tt.args.h1, tt.args.h2, tt.args.h3)
			assert.Equal(t, tt.want, got)
		})
	}
}
