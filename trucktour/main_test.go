package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_truckTour(t *testing.T) {
	type args struct {
		petrolpumps [][]int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "test with empty array",
			args: args{
				petrolpumps: [][]int32{},
			},
			want: -1,
		},
		{
			name: "test with valid array",
			args: args{
				petrolpumps: [][]int32{{1, 5}, {10, 3}, {3, 4}},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := truckTour(tt.args.petrolpumps)
			assert.Equal(t, tt.want, got)
		})
	}
}
