package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hackerlandRadioTransmitters(t *testing.T) {
	type args struct {
		x []int
		k int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "test with empty array",
			args: args{
				x: []int{},
				k: 0,
			},
			want: 0,
		},
		{
			name: "test with valid array",
			args: args{
				x: []int{1, 2, 3, 4, 5},
				k: 1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hackerlandRadioTransmitters(tt.args.x, tt.args.k)
			assert.Equal(t, tt.want, got)
		})
	}
}
