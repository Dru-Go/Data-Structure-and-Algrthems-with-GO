package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_gridChallenge(t *testing.T) {
	type args struct {
		grid []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with empty array",
			args: args{
				grid: []string{},
			},
			want: "NO",
		},
		{
			name: "Test with empty array",
			args: args{
				grid: []string{"ebacd", "fghij", "olmkn", "trpqs", "xywuv"},
			},
			want: "YES",
		},
		{
			name: "Test with empty array",
			args: args{
				grid: []string{"kc", "iu", "3", "uxf", "vof", "hmp"},
			},
			want: "YES",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gridChallenge(tt.args.grid)
			assert.Equal(t, tt.want, got)
		})
	}
}
