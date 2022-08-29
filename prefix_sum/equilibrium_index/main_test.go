package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_equilibriumIndex(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test with empty array",
			args: args{
				arr: []int32{},
			},
			want: -1,
		},
		{
			name: "test with array",
			args: args{
				arr: []int32{1, 2, 3, 2, 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := equilibriumIndex(tt.args.arr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_efficientEquilibriumIndex(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test with empty array",
			args: args{
				arr: []int32{},
			},
			want: -1,
		},
		{
			name: "test with array",
			args: args{
				arr: []int32{1, 2, 3, 2, 1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := efficientEquilibriumIndex(tt.args.arr)
			assert.Equal(t, tt.want, got)
		})
	}
}
