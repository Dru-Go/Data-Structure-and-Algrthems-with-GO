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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := balancedSums(tt.args.arr)
			assert.Equal(t, tt.want, got)
		})
	}
}
