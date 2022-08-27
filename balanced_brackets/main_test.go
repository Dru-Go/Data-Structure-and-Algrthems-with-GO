package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isBalanced(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test with empty string",
			args: args{
				s: "",
			},
			want: "NO",
		},
		{
			name: "test with string",
			args: args{
				s: "[()]",
			},
			want: "YES",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isBalanced(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
