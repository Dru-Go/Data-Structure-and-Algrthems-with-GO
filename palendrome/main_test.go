package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_palindromeIndex(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "Test with valid string",
			args: args{
				s: "aaa",
			},
			want: -1,
		},
		{
			name: "Test with valid string",
			args: args{
				s: "quyjjdcgsvvsgcdjjyq",
			},
			want: 1,
		},
		{
			name: "Test with valid string",
			args: args{
				s: "lhrxvssvxrhl",
			},
			want: -1,
		},
		{
			name: "Test with valid string",
			args: args{
				s: "hgygsvlfcwnswtuhmyaljkqlqjjqlqkjlaymhutwsnwcwflvsgygh",
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := palindromeIndex(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
