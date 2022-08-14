package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_counterGame(t *testing.T) {
	type args struct {
		n int64
	}
	// var Louise = "Louise"
	var Richard = "Richard"
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "testing with 0",
			args: args{
				n: 0,
			},
			want: Richard,
		},
		{
			name: "testing with non zero",
			args: args{
				n: 6,
			},
			want: Richard,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := counterGame(tt.args.n)
			assert.Equal(t, tt.want, got)
		})
	}
}
