package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxMin(t *testing.T) {
	type args struct {
		k   int32
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "test empty array",
			args: args{
				k:   4,
				arr: []int32{},
			},
			want: 0,
		},
		{
			name: "test with number array",
			args: args{
				k:   4,
				arr: []int32{1, 2, 3, 4, 10, 20, 30, 40, 100, 200},
			},
			want: 3,
		},
		{
			name: "test with number array",
			args: args{
				k:   5,
				arr: []int32{4504, 1520, 5857, 4094, 4157, 3902, 822, 6643, 2422, 7288, 8245, 9948, 2822, 1784, 7802, 3142, 9739, 5629},
			},
			want: 1335,
		},
		{
			name: "test with number array",
			args: args{
				k:   3,
				arr: []int32{100, 200, 300, 350, 400, 401, 402},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxMin(tt.args.k, tt.args.arr)
			assert.Equal(t, tt.want, got)
		})
	}
}
