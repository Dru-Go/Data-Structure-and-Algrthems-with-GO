package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sumXor(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Testing with n equal to 0",
			args: args{
				n: 0,
			},
			want: 1,
		},
		{
			name: "Testing with different n values",
			args: args{
				n: 5,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sumXor(tt.args.n)
			assert.Equal(t, tt.want, got)
		})
	}
}
