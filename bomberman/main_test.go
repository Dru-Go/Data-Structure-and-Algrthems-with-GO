package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_bomberMan(t *testing.T) {
	type args struct {
		n    int32
		grid []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test with empty grid",
			args: args{
				n:    1,
				grid: []string{},
			},
			want: []string{},
		},
		{
			name: "Test with some grid and n equal 2",
			args: args{
				n:    2,
				grid: []string{"000", "..."},
			},
			want: []string{"000", "000"},
		},
		{
			name: "Test with some grid and n equal 3 #1",
			args: args{
				n:    3,
				grid: []string{"000", "..."},
			},
			want: []string{"000", "000"},
		},
		{
			name: "Test with some grid and n equal 3 #2",
			args: args{
				n:    3,
				grid: []string{".......", "...0...", "....0..", "...0...", ".......", "00.....", "00....."},
			},
			want: []string{"000.000", "00...00", "000...0", "..00.00", "...0000", "...0000"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bomberMan(tt.args.n, tt.args.grid)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_permutationIndexOfRune(t *testing.T) {
	type args struct {
		str       string
		subString byte
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test with empty string",
			args: args{
				str:       "",
				subString: byte('0'),
			},
			want: []int{},
		},
		{
			name: "test with string",
			args: args{
				str:       "010",
				subString: byte('0'),
			},
			want: []int{0, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := indexOfRune(tt.args.str, tt.args.subString)
			assert.Equal(t, tt.want, got)
		})
	}
}
