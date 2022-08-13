package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_balancedSums(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// {
		// 	name: "testing with empty array",
		// 	args: args{
		// 		arr: []int32{},
		// 	},
		// 	want: "NO",
		// },
		// {
		// 	name: "testing with array",
		// 	args: args{
		// 		arr: []int32{1, 2, 3},
		// 	},
		// 	want: "NO",
		// },
		// {
		// 	name: "testing with array",
		// 	args: args{
		// 		arr: []int32{1, 2, 3, 3},
		// 	},
		// 	want: "YES",
		// },
		// {
		// 	name: "testing with array",
		// 	args: args{
		// 		arr: []int32{2, 0, 0, 0},
		// 	},
		// 	want: "YES",
		// },
		{
			name: "testing with array",
			args: args{
				arr: []int32{1},
			},
			want: "YES",
		},
		// {
		// 	name: "testing with array",
		// 	args: args{
		// 		arr: []int32{2},
		// 	},
		// 	want: "NO",
		// },
		// {
		// 	name: "testing with array",
		// 	args: args{
		// 		arr: []int32{1, 2},
		// 	},
		// 	want: "NO",
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := balancedSums(tt.args.arr)
			assert.Equal(t, tt.want, got)
		})
	}
}
