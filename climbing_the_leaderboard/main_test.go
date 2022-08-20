package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_climbingLeaderboard(t *testing.T) {
	type args struct {
		ranked []int32
		player []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "Test with empty array",
			args: args{
				ranked: []int32{},
				player: []int32{},
			},
			want: []int32{},
		},
		{
			name: "Test with valid array",
			args: args{
				ranked: []int32{100, 90, 90, 80},
				player: []int32{70, 80, 105},
			},
			want: []int32{4, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := climbingLeaderboard(tt.args.ranked, tt.args.player)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		s []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "Test with empty array",
			args: args{
				s: []int32{},
			},
			want: []int32{},
		},
		{
			name: "Test with valid array",
			args: args{
				s: []int32{1, 1, 1, 2},
			},
			want: []int32{1, 2},
		},
		{
			name: "Test with unsorted array",
			args: args{
				s: []int32{2, 1, 1, 1, 2},
			},
			want: []int32{2, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
