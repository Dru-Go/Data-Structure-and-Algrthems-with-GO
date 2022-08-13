package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoArrays(t *testing.T) {
	type args struct {
		k int32
		A []int32
		B []int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "right permutation",
			args: args{
				k: 10,
				A: []int32{2, 1, 3},
				B: []int32{7, 8, 9},
			},
			want: "YES",
		},
		{
			name: "wrong permutation",
			args: args{
				k: 5,
				A: []int32{1, 2, 2, 1},
				B: []int32{3, 3, 3, 4},
			},
			want: "NO",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoArrays(tt.args.k, tt.args.A, tt.args.B)
			assert.Equal(t, tt.want, got)
		})
	}
}
