package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_birthday(t *testing.T) {
	type args struct {
		s []int32
		d int32
		m int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "testing with incorrect result 1",
			args: args{
				s: []int32{},
				d: 5,
				m: 2,
			},
			want: 0,
		},
		{
			name: "testing with incorrect result 2",
			args: args{
				s: []int32{1, 1, 1, 1, 1, 1},
				d: 3,
				m: 2,
			},
			want: 0,
		},
		{
			name: "testing with incorrect result 2",
			args: args{
				s: []int32{1, 2, 1, 3, 2},
				d: 3,
				m: 2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := birthday(tt.args.s, tt.args.d, tt.args.m)
			assert.Equal(t, got, tt.want)
		})
	}
}
