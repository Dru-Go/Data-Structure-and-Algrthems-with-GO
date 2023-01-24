package chunk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunk(t *testing.T) {
	type args struct {
		collection []int16
		size       int
	}
	tests := []struct {
		name string
		args args
		want [][]int16
	}{
		{
			name: "Test with empty collection",
			args: args{
				collection: []int16{},
				size:       10,
			},
			want: [][]int16{},
		},
		{
			name: "Test with collection",
			args: args{
				collection: []int16{1, 2, 3, 4, 5, 6},
				size:       2,
			},
			want: [][]int16{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			name: "Test with no size",
			args: args{
				collection: []int16{1, 2, 3, 4, 5, 6},
				size:       0,
			},
			want: [][]int16{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Chunk(tt.args.collection, tt.args.size)
			assert.Equal(t, tt.want, got)
		})
	}
}
