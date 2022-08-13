package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pageCount(t *testing.T) {
	type args struct {
		n int32
		p int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "test with no pages",
			args: args{
				n: 0,
				p: 5,
			},
			want: 0,
		},
		{
			name: "test with pages",
			args: args{
				n: 11,
				p: 7,
			},
			want: 2,
		},
		{
			name: "test with pages",
			args: args{
				n: 6,
				p: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pageCount(tt.args.n, tt.args.p)
			assert.Equal(t, tt.want, got)
		})
	}
}
