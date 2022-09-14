package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_noPrefix(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test with empty array",
			args: args{
				words: []string{},
			},
			want: "BAD SET",
		},
		{
			name: "test with a valid array",
			args: args{
				words: []string{"abebe", "abebu"},
			},
			want: "BAD SET",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := noPrefix(tt.args.words)
			assert.Equal(t, tt.want, got)
		})
	}
}
