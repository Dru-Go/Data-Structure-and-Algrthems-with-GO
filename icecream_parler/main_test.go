package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_icecreamParlor(t *testing.T) {
	type args struct {
		m   int32
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "Test with empty array",
			args: args{
				m:   4,
				arr: []int32{},
			},
			want: []int32{},
		},
		{
			name: "Test with a valid array #1",
			args: args{
				m:   5,
				arr: []int32{1, 4, 5, 3, 2},
			},
			want: []int32{1, 4},
		},
		{
			name: "Test with a valid array #2",
			args: args{
				m:   4,
				arr: []int32{2, 2, 4, 3},
			},
			want: []int32{1, 2},
		},
		{
			name: "Test with a valid array #3",
			args: args{
				m:   4,
				arr: []int32{10, 1, 4, 5, 3, 2, 8},
			},
			want: []int32{1, 3},
		},
		{
			name: "Test with a valid array #4",
			args: args{
				m:   6,
				arr: []int32{1, 3, 4, 6, 7, 9},
			},
			want: []int32{2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := icecreamParlor(tt.args.m, tt.args.arr)
			assert.Equal(t, tt.want, got)
		})
	}
}
