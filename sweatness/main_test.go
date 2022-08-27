package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_cookies(t *testing.T) {
	type args struct {
		k int32
		A []int
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "Test with empty array",
			args: args{
				k: 5,
				A: []int{},
			},
			want: -1,
		},
		{
			name: "Test with  array",
			args: args{
				k: 7,
				A: []int{1, 2, 3, 9, 10, 12},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cookies(tt.args.k, tt.args.A)
			assert.Equal(t, tt.want, got)
		})
	}
}
