package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNPF(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test with zero",
			args: args{n: 0},
			want: 0,
		},
		{
			name: "test with non zero digit",
			args: args{n: 100},
			want: 1060,
		},
		{
			name: "test with test digits #1",
			args: args{n: 4},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NPF(tt.args.n)
			assert.Equal(t, tt.want, got)
		})
	}
}
