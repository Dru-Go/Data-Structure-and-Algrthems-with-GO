package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumBribes(t *testing.T) {
	type args struct {
		q []int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test with empty array",
			args: args{
				q: []int32{},
			},
			want: "0",
		},
		{
			name: "test with mixed array",
			args: args{
				q: []int32{3, 1},
			},
			want: "1",
		},
		{
			name: "test with mixed array",
			args: args{
				q: []int32{1, 3},
			},
			want: "0",
		},
		{
			name: "test with mixed array",
			args: args{
				q: []int32{1, 5, 4, 3},
			},
			want: "2",
		},
		{
			name: "test with mixed array",
			args: args{
				q: []int32{2, 1, 5, 3, 4},
			},
			want: "3",
		},
		{
			name: "test with mixed array",
			args: args{
				q: []int32{2, 5, 1, 3, 4},
			},
			want: "Too chaotic",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumBribes(tt.args.q)
			assert.Equal(t, tt.want, got)
		})
	}
}
