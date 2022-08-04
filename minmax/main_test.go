package main

import (
	"reflect"
	"testing"
)

func Test_miniMaxSum(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "Testing min max",
			args: args{
				arr: []int32{1, 2, 3, 4, 5},
			},
			want: []int64{10, 14},
		},
		{
			name: "Testing min max with one element",
			args: args{
				arr: []int32{1},
			},
			want: []int64{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := miniMaxSum(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("miniMaxSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
