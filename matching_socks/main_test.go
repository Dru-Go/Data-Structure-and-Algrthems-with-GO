package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sockMerchant(t *testing.T) {
	type args struct {
		n  int32
		ar []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "Testing with empty array",
			args: args{
				n:  0,
				ar: []int32{},
			},
			want: 0,
		},
		{
			name: "Testing with empty array",
			args: args{
				n:  9,
				ar: []int32{10, 20, 20, 10, 10, 30, 50, 10, 20},
			},
			want: 3,
		},
		{
			name: "Testing with empty array",
			args: args{
				n:  9,
				ar: []int32{1, 1, 3, 1, 2, 1, 3, 3, 3, 3},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sockMerchant(tt.args.n, tt.args.ar)
			assert.Equal(t, tt.want, got)
		})
	}
}
