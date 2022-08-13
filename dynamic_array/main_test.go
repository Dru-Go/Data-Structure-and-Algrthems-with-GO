package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_dynamicArray(t *testing.T) {
	type args struct {
		n       int32
		queries [][]int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dynamicArray(tt.args.n, tt.args.queries)
			assert.Equal(t, tt.want, got)
		})
	}
}
