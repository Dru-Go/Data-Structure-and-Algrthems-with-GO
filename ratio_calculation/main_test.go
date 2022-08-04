package main

import (
	"reflect"
	"testing"
)

func Test_plusMinus1(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testing the output",
			args: args{arr: []int32{-4, 3, -9, 0, 4, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plusMinus(tt.args.arr)
		})
	}
}

func Test_plusMinus(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "testing the output",
			args: args{arr: []int32{-4, 3, -9, 0, 4, 1}},
			want: []string{
				"0.500000",
				"0.166667",
				"0.333333",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusMinus(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusMinus() = %v, want %v", got, tt.want)
			}
		})
	}
}
