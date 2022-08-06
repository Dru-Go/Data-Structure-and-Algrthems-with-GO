package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMedian(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMedian(tt.args.arr)
			assert.Equal(t, got, tt.want)
		})
	}
}
