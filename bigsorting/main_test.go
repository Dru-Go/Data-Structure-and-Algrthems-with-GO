package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_bigSorting(t *testing.T) {
	type args struct {
		unsorted []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test with empty array",
			args: args{
				unsorted: []string{},
			},
			want: []string{},
		},
		{
			name: "test with valid array",
			args: args{
				unsorted: []string{"100", "200", "120", "100", "110"},
			},
			want: []string{"100", "100", "110", "120", "200"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bigSorting(tt.args.unsorted)
			assert.Equal(t, tt.want, got)
		})
	}
}
