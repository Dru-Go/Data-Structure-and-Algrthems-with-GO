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
		name  string
		args  args
		want  string
		want1 string
	}{
		{
			name: "test with an empty array",
			args: args{
				words: []string{},
			},
			want:  "GOOD SET",
			want1: "",
		},
		{
			name: "test with a valid string array",
			args: args{
				words: []string{"aab", "defgab", "abcde", "aabcde", "bbbbbbbbbb", "jabjjjad"},
			},
			want:  "BAD SET",
			want1: "aabcde",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := noPrefix(tt.args.words)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)

		})
	}
}
