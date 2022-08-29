package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_binarySearch(t *testing.T) {
	type args struct {
		arr    []int32
		target int32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test with empty array",
			args: args{
				arr:    []int32{},
				target: 0,
			},
			want: -1,
		},
		{
			name: "test with array",
			args: args{
				arr:    []int32{1, 2, 5, 6, 7},
				target: 6,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := binarySearch(tt.args.arr, tt.args.target)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_recursiveBinarySearch(t *testing.T) {
	type args struct {
		array     []int32
		target    int32
		lowIndex  int
		highIndex int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test with empty array",
			args: args{
				array:     []int32{},
				target:    0,
				lowIndex:  0,
				highIndex: 0,
			},
			want: -1,
		},
		{
			name: "test with array",
			args: args{
				array:     []int32{1, 2, 5, 6, 7},
				target:    6,
				lowIndex:  0,
				highIndex: len([]int32{1, 2, 5, 6, 7}),
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := recursiveBinarySearch(tt.args.array, tt.args.target, tt.args.lowIndex, tt.args.highIndex)
			assert.Equal(t, tt.want, got)
		})
	}
}
